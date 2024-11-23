package entities

type GetVotingResultsResponse struct {
	StudentID          string `json:"student_id"`
	CandidateFirstName string `json:"candidate_first_name"`
	CandidateLastName  string `json:"candidate_last_name"`
	CandidatePosition  string `json:"candidate_position"`
	TransactionID      string `json:"transaction_id"`
	IsValid            bool   `json:"is_valid"`
}

type VoteDetails struct {
	StudentID    string `json:"student_id"`
	VoterAddress string `json:"voter_address"`
	CandidateID  string `json:"candidate_id"`
}

type CastVoteRequest struct {
	StudentID   string `json:"student_id" validate:"required"`
	CandidateID string `json:"candidate_id" validate:"required"`
}

type VoteDetailsResponse struct {
	StudentID    string `json:"student_id"`
	VoterAddress string `json:"voter_address"`
	CandidateID  string `json:"candidate_id"`
}

type EndVotingRequest struct {
	AdminAddress string `json:"admin_address" validate:"required"`
}

type GetCandidatesRankingResponse struct {
	CandidateID string `json:"candidate_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Position    string `json:"position"`
	Votes       int    `json:"votes"`
}
