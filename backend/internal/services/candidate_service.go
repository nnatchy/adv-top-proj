package services

import (
	"context"

	"github.com/nnatchy/adv-top-proj/internal/entities"
	"github.com/nnatchy/adv-top-proj/internal/infrastructure/log"
	"github.com/nnatchy/adv-top-proj/internal/repositories"
)

// ICandidateService defines the interface for the candidate service
type ICandidateService interface {
	SearchCandidates(ctx context.Context, position string) ([]entities.SearchCandidatesResponse, error)
}

// CandidateService is the implementation of ICandidateService
type CandidateService struct {
	repo repositories.RepositoryGateway
	log  *log.Logger
}

// NewCandidateService creates a new instance of CandidateService
func NewCandidateService(
	repo repositories.RepositoryGateway,
	log *log.Logger,
) ICandidateService {
	return &CandidateService{
		repo: repo,
		log:  log,
	}
}

// SearchCandidates fetches and formats candidate data based on position
func (s *CandidateService) SearchCandidates(ctx context.Context, position string) ([]entities.SearchCandidatesResponse, error) {
	// Fetch candidates from the repository
	s.log.InfoWithID(ctx, "[Service: SearchCandidates] Called ")
	candidates, err := s.repo.CandidateRepository.SearchCandidates(ctx, position)
	if err != nil {
		s.log.ErrorWithID(ctx, "[Service: SearchCandidates] Failed to fetch candidates", err)
		return nil, err
	}

	// Format response
	var response []entities.SearchCandidatesResponse
	for _, candidate := range candidates {
		response = append(response, entities.SearchCandidatesResponse{
			StudentID:          candidate.StudentID,
			CandidateFirstName: candidate.FirstName,
			CandidateLastName:  candidate.LastName,
			CandidatePosition:  candidate.Position,
			CandidateMajor:     candidate.Major,
			CandidateID:        candidate.ID.String(),
		})
	}

	return response, nil
}
