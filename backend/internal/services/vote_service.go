package services

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/nnatchy/adv-top-proj/internal/entities"
	"github.com/nnatchy/adv-top-proj/internal/infrastructure/ethereum/solidity/voting"
	"github.com/nnatchy/adv-top-proj/internal/infrastructure/log"
	"github.com/nnatchy/adv-top-proj/internal/models"
	"github.com/nnatchy/adv-top-proj/internal/repositories"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/dig"
)

type IVoteService interface {
	CastVote(ctx context.Context, req entities.CastVoteRequest) (string, error)
	EndVoting(ctx context.Context, req entities.EndVotingRequest) (string, error)
	GetVoteDetails(ctx context.Context, studentID string) (*entities.VoteDetailsResponse, error)
	GetVotingResults(ctx context.Context) ([]entities.GetVotingResultsResponse, error)
	GetAllVotesFromEvents(ctx context.Context) ([]entities.VoteDetailsResponse, error)
	ValidateVote(ctx context.Context, vote *models.Vote) (bool, error)
	GetCandidatesRanking(ctx context.Context) ([]entities.GetCandidatesRankingResponse, error)
}

type VoteService struct {
	repo           repositories.RepositoryGateway
	log            *log.Logger
	votingContract *voting.Voting
	transactor     *bind.TransactOpts
}

type VoteServiceParams struct {
	dig.In
	Repo           repositories.RepositoryGateway
	Log            *log.Logger
	VotingContract *voting.Voting
	Transactor     *bind.TransactOpts `name:"VotingTransactor"`
}

func NewVoteService(params VoteServiceParams) IVoteService {
	return &VoteService{
		repo:           params.Repo,
		log:            params.Log,
		votingContract: params.VotingContract,
		transactor:     params.Transactor,
	}
}

// CastVote handles casting a vote
func (s *VoteService) CastVote(ctx context.Context, req entities.CastVoteRequest) (string, error) {
	s.log.InfoWithID(ctx, "[Service: CastVote] Initiating vote casting")

	// Fetch user wallet address
	user, err := s.repo.UserRepository.FindByStudentId(ctx, req.StudentID)
	if err != nil || user == nil || user.WalletAddress == "" {
		s.log.ErrorWithID(ctx, "[Service: CastVote] Failed to fetch wallet address", err)
		return "", fmt.Errorf("failed to fetch wallet address: %w", err)
	}

	s.log.InfoWithID(ctx, "[Service: CastVote] User retrieved", "WalletAddress", user.WalletAddress)

	// Check if the student has already voted
	// hasVoted, err := s.votingContract.HasVoted(&bind.CallOpts{Context: ctx, From: s.transactor.From}, req.StudentID)
	// if err != nil {
	// 	s.log.ErrorWithID(ctx, "[Service: CastVote] Error checking voting status", err)
	// 	return "", fmt.Errorf("error checking voting status: %w", err)
	// }

	// if hasVoted {
	// 	s.log.InfoWithID(ctx, "[Service: CastVote] Student has already voted", "studentID", req.StudentID)
	// 	return "", fmt.Errorf("student has already voted")
	// }

	// Create a vote object for validation and tracking
	vote := &models.Vote{
		StudentID:     req.StudentID,
		VoterAddress:  user.WalletAddress,
		CandidateID:   req.CandidateID,
		TransactionID: "",
		IsValid:       false, // Initially marked as invalid
		CreatedAt:     time.Now(),
	}

	// Validate the vote details
	isValid, err := s.ValidateVote(ctx, vote)
	if err != nil {
		s.log.ErrorWithID(ctx, "[Service: CastVote] Vote validation failed", err)
		return "", fmt.Errorf("vote validation failed: %w", err)
	}

	if !isValid {
		s.log.InfoWithID(ctx, "[Service: CastVote] Invalid vote detected", "studentID", req.StudentID)
		return "", fmt.Errorf("invalid vote")
	}
	// Ensure vote choice is valid
	candidate, err := s.repo.CandidateRepository.GetCandidateById(ctx, vote.CandidateID)
	if err != nil || candidate == nil {
		s.log.ErrorWithID(ctx, "[Service: validateVote] Invalid candidate", err)
		return "", fmt.Errorf("invalid candidate for CandidateID: %s", vote.CandidateID)
	}

	s.log.InfoWithID(ctx, "[Service: CastVote] Vote is valid", "StudentID", req.StudentID)

	// Cast the vote on-chain
	tx, err := s.votingContract.CastVote(s.transactor, req.StudentID, common.HexToAddress(user.WalletAddress), req.CandidateID)
	if err != nil {
		s.log.ErrorWithID(ctx, "[Service: CastVote] Error casting vote on blockchain", err)
		return "", fmt.Errorf("blockchain transaction failed: %w", err)
	}

	txHash := tx.Hash().Hex()
	s.log.InfoWithID(ctx, "[Service: CastVote] Vote cast successfully", "Transaction Hash", txHash)

	// Update the vote with transaction details and mark as valid
	vote.TransactionID = txHash
	vote.IsValid = true

	if user.HasVoteYear {
		s.log.InfoWithID(ctx, "[Service: CastVote] Student has already voted by year", "studentID", req.StudentID)
	}

	if user.HasVoteRep {
		s.log.InfoWithID(ctx, "[Service: CastVote] Student has already voted by year", "studentID", req.StudentID)
	}

	if user.HasVoteYear && (candidate.Position == "year_one" || candidate.Position == "year_two" || candidate.Position == "year_three" || candidate.Position == "year_four") {
		s.log.InfoWithID(ctx, "[Service: CastVote] Student has already voted by year", "studentID", req.StudentID)
		vote.IsValid = false
	}

	if user.HasVoteRep && candidate.Position == "representative" {
		s.log.InfoWithID(ctx, "[Service: CastVote] Student has already votedby rep", "studentID", req.StudentID)
		vote.IsValid = false
	}

	// Save the vote to the database
	if err := s.repo.VoteRepository.SaveVote(ctx, vote); err != nil {
		s.log.ErrorWithID(ctx, "[Service: CastVote] Error saving vote to database", err)
		return txHash, fmt.Errorf("failed to save vote to database: %w", err)
	}

	s.log.InfoWithID(ctx, "[Service: CastVote] Vote saved successfully", "StudentID", req.StudentID)
	return txHash, nil
}

// validateVote checks the validity of a vote before casting
func (s *VoteService) ValidateVote(ctx context.Context, vote *models.Vote) (bool, error) {
	s.log.InfoWithID(ctx, "[Service: validateVote] Validating vote", "StudentID", vote.StudentID)

	// Check if the voter's wallet matches the student ID
	user, err := s.repo.UserRepository.FindByStudentId(ctx, vote.StudentID)
	if err != nil || user == nil {
		s.log.ErrorWithID(ctx, "[Service: validateVote] User not found", err)
		return false, fmt.Errorf("user not found for StudentID: %s", vote.StudentID)
	}

	if user.WalletAddress != vote.VoterAddress {
		s.log.ErrorWithID(ctx, "[Service: validateVote] Wallet mismatch", "Expected", user.WalletAddress, "Found", vote.VoterAddress)
		return false, fmt.Errorf("wallet address does not match for StudentID: %s", vote.StudentID)
	}

	// // Check if the vote already exists in the database
	// existingVote, err := s.repo.VoteRepository.FindVoteByStudentID(ctx, vote.StudentID)
	// if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
	// 	s.log.ErrorWithID(ctx, "[Service: validateVote] Error fetching existing vote", err)
	// 	return false, fmt.Errorf("error fetching existing vote: %w", err)
	// }

	// if existingVote != nil && existingVote.IsValid {
	// 	s.log.InfoWithID(ctx, "[Service: validateVote] Vote already exists and is valid", "StudentID", vote.StudentID)
	// 	return false, fmt.Errorf("vote already exists for StudentID: %s", vote.StudentID)
	// }

	s.log.InfoWithID(ctx, "[Service: validateVote] Vote validation successful", "StudentID", vote.StudentID, "CandidateID", vote.CandidateID)
	return true, nil
}

// EndVoting ends the voting process
func (s *VoteService) EndVoting(ctx context.Context, req entities.EndVotingRequest) (string, error) {
	s.log.InfoWithID(ctx, "[Service: EndVoting] Ending voting process")

	// Use the transactor provided via DI
	auth := s.transactor
	auth.Context = ctx

	// Check if the caller is the owner
	if auth.From.Hex() != req.AdminAddress {
		s.log.ErrorWithID(ctx, "[Service: EndVoting] Unauthorized access", "AdminAddress", req.AdminAddress)
		return "", fmt.Errorf("only the owner can end the voting")
	}

	tx, err := s.votingContract.EndVoting(auth)
	if err != nil {
		s.log.ErrorWithID(ctx, "[Service: EndVoting] Error ending voting on blockchain", err)
		return "", fmt.Errorf("blockchain transaction failed: %w", err)
	}

	txHash := tx.Hash().Hex()
	s.log.InfoWithID(ctx, "[Service: EndVoting] Voting ended successfully", "Transaction Hash", txHash)
	return txHash, nil
}

// GetVoteDetails retrieves the details of a vote
func (s *VoteService) GetVoteDetails(ctx context.Context, studentID string) (*entities.VoteDetailsResponse, error) {
	s.log.InfoWithID(ctx, "[Service: GetVoteDetails] Retrieving vote details", "studentID", studentID)

	// Fetch vote from the database
	vote, err := s.repo.VoteRepository.FindVoteByStudentID(ctx, studentID)
	if err != nil {
		s.log.ErrorWithID(ctx, "[Service: GetVoteDetails] Error fetching vote from database", err)
		return nil, fmt.Errorf("failed to fetch vote from database: %w", err)
	}

	if vote == nil {
		s.log.ErrorWithID(ctx, "[Service: GetVoteDetails] No vote found for StudentID", "studentID", studentID)
		return nil, fmt.Errorf("no vote found for studentID: %s", studentID)
	}

	// Prepare response
	voteDetails := &entities.VoteDetailsResponse{
		StudentID:    vote.StudentID,
		VoterAddress: vote.VoterAddress,
		CandidateID:  vote.CandidateID,
	}

	return voteDetails, nil
}

// GetVotingResults retrieves all voting results from the database
func (s *VoteService) GetVotingResults(ctx context.Context) ([]entities.GetVotingResultsResponse, error) {
	s.log.InfoWithID(ctx, "[Service: GetVotingResults] Fetching voting results")

	// Fetch all votes from the database
	votes, err := s.repo.VoteRepository.GetAllVotes(ctx)
	if err != nil {
		s.log.ErrorWithID(ctx, "[Service: GetVotingResults] Failed to fetch votes from database", err)
		return nil, fmt.Errorf("failed to fetch votes from database: %w", err)
	}

	if len(votes) == 0 {
		s.log.InfoWithID(ctx, "[Service: GetVotingResults] No votes found in the database")
		return nil, nil
	}

	results := make([]entities.GetVotingResultsResponse, 0, len(votes))
	for _, vote := range votes {
		candidate, err := s.repo.CandidateRepository.GetCandidateById(ctx, vote.CandidateID)
		if err != nil {
			// Log the error and continue with other votes
			s.log.ErrorWithID(ctx, "[Service: GetVotingResults] Error fetching candidate details", "CandidateID", vote.CandidateID, err)
			continue
		}

		results = append(results, entities.GetVotingResultsResponse{
			StudentID:          vote.StudentID,
			CandidateFirstName: candidate.FirstName,
			CandidateLastName:  candidate.LastName,
			CandidatePosition:  candidate.Position,
			TransactionID:      vote.TransactionID,
			IsValid:            vote.IsValid,
		})
	}

	s.log.InfoWithID(ctx, "[Service: GetVotingResults] Successfully fetched voting results", "Count", len(results))
	return results, nil
}

// GetAllVotesFromEvents retrieves all votes by filtering the VoteCast events
func (s *VoteService) GetAllVotesFromEvents(ctx context.Context) ([]entities.VoteDetailsResponse, error) {
	s.log.InfoWithID(ctx, "[Service: GetAllVotesFromEvents] Fetching all votes from events")

	// Create FilterOpts with context
	filterOpts := &bind.FilterOpts{
		Context: ctx,
		Start:   0, // Adjust if you know the deployment block
		End:     nil,
	}

	// Filter the events
	iter, err := s.votingContract.FilterVoteCast(filterOpts)
	if err != nil {
		s.log.ErrorWithID(ctx, "[Service: GetAllVotesFromEvents] Error filtering VoteCast events", err)
		return nil, fmt.Errorf("error filtering VoteCast events: %w", err)
	}
	defer iter.Close()

	var voteDetailsList []entities.VoteDetailsResponse
	count := 0 // Counter for events
	for iter.Next() {
		event := iter.Event

		s.log.InfoWithID(ctx, "[Service: GetAllVotesFromEvents] Found VoteCast event", "StudentID", event.StudentId, "VoterAddress", event.VoterAddress.Hex(), "VoteChoice", event.VoteChoice)

		voteDetails := entities.VoteDetailsResponse{
			StudentID:    event.StudentId,
			VoterAddress: event.VoterAddress.Hex(),
			CandidateID:  event.VoteChoice,
		}
		voteDetailsList = append(voteDetailsList, voteDetails)
		count++
	}

	s.log.InfoWithID(ctx, "[Service: GetAllVotesFromEvents] Total events processed", "Count", count)

	if err := iter.Error(); err != nil {
		s.log.ErrorWithID(ctx, "[Service: GetAllVotesFromEvents] Iterator error", err)
		return nil, fmt.Errorf("iterator error: %w", err)
	}

	return voteDetailsList, nil
}

func (s *VoteService) GetCandidatesRanking(ctx context.Context) ([]entities.GetCandidatesRankingResponse, error) {
	s.log.InfoWithID(ctx, "[Service: GetCandidatesRanking] Called")

	// Step 1: Retrieve all votes
	votes, err := s.repo.VoteRepository.GetAllVotes(ctx)
	if err != nil {
		s.log.ErrorWithID(ctx, "[Service: GetCandidatesRanking] Failed to fetch votes", err)
		return nil, fmt.Errorf("failed to fetch votes: %w", err)
	}

	// Step 2: Aggregate votes by CandidateID
	voteCounts := make(map[string]int) // CandidateID -> Vote Count
	for _, vote := range votes {
		if vote.IsValid {
			voteCounts[vote.CandidateID]++
		}
	}

	// Step 3: Fetch candidate details for all unique CandidateIDs
	candidateIDs := make([]string, 0, len(voteCounts))
	for candidateID := range voteCounts {
		candidateIDs = append(candidateIDs, candidateID)
	}

	candidates, err := s.repo.CandidateRepository.GetCandidateByIds(ctx, candidateIDs)
	if err != nil {
		s.log.ErrorWithID(ctx, "[Service: GetCandidatesRanking] Failed to fetch candidate details", err)
		return nil, fmt.Errorf("failed to fetch candidate details: %w", err)
	}

	s.log.InfoWithID(ctx, "[Service: GetCandidatesRanking] Candidates ", candidates)

	// Step 4: Map candidates by ID for quick lookup
	candidateMap := make(map[string]models.Candidate)
	for _, candidate := range candidates {
		candidateMap[candidate.ID.String()] = *candidate
	}

	// Step 5: Build rankings grouped by position
	positionRankings := make(map[string][]entities.GetCandidatesRankingResponse) // Position -> List of candidates
	for candidateID, count := range voteCounts {
		candidate, exists := candidateMap[candidateID]
		if !exists {
			s.log.ErrorWithID(ctx, "[Service: GetCandidatesRanking] Candidate details missing for ID: ", candidateID)
			continue // Skip if candidate details are missing
		}

		voteDetails := entities.GetCandidatesRankingResponse{
			CandidateID: candidate.ID.String(),
			FirstName:   candidate.FirstName,
			LastName:    candidate.LastName,
			Position:    candidate.Position,
			Votes:       count,
		}

		positionRankings[candidate.Position] = append(positionRankings[candidate.Position], voteDetails)
	}

	// Step 6: Sort candidates by votes within each position
	var rankedResults []entities.GetCandidatesRankingResponse
	for position, candidates := range positionRankings {
		// Sort candidates by descending vote count
		sort.Slice(candidates, func(i, j int) bool {
			return candidates[i].Votes > candidates[j].Votes
		})

		// Log the ranking for this position
		s.log.InfoWithID(ctx, "[Service: GetCandidatesRanking] Ranking for position: ", position, candidates)

		// Add sorted candidates to the final results
		rankedResults = append(rankedResults, candidates...)
	}

	s.log.InfoWithID(ctx, "[Service: GetCandidatesRanking] Successfully retrieved candidate rankings")
	return rankedResults, nil
}
