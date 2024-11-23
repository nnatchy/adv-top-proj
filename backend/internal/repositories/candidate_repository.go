package repositories

import (
	"context"

	"github.com/nnatchy/adv-top-proj/internal/infrastructure/log"
	"github.com/nnatchy/adv-top-proj/internal/models"
	"gorm.io/gorm"
)

type ICandidateRepository interface {
	GetCandidateById(ctx context.Context, candidateID string) (*models.Candidate, error)
	SearchCandidates(ctx context.Context, position string) ([]*models.Candidate, error)
	GetCandidates(ctx context.Context) ([]*models.Candidate, error)
	GetCandidateByIds(ctx context.Context, candidateIDs []string) ([]*models.Candidate, error)
}

type CandidateRepository struct {
	db  *gorm.DB
	log *log.Logger
}

func NewCandidateRepository(db *gorm.DB, log *log.Logger) ICandidateRepository {
	return &CandidateRepository{
		db:  db,
		log: log,
	}
}

func (r *CandidateRepository) GetCandidateById(ctx context.Context, candidateID string) (*models.Candidate, error) {
	r.log.InfoWithID(ctx, "[Repository: GetCandidateById] Called")
	var candidate models.Candidate
	if err := r.db.WithContext(ctx).Where("id = ?", candidateID).First(&candidate).Error; err != nil {
		r.log.ErrorWithID(ctx, "[Repository: GetCandidateById] Error Querying Candidate", err)
		return nil, err
	}
	return &candidate, nil
}

func (r *CandidateRepository) SearchCandidates(ctx context.Context, position string) ([]*models.Candidate, error) {
	r.log.InfoWithID(ctx, "[Repository: SearchCandidates] Called")
	var candidates []*models.Candidate
	if err := r.db.WithContext(ctx).Where("position = ?", position).Find(&candidates).Error; err != nil {
		r.log.ErrorWithID(ctx, "[Repository: SearchCandidates] Error Querying Candidates", err)
		return nil, err
	}
	return candidates, nil
}

func (r *CandidateRepository) GetCandidates(ctx context.Context) ([]*models.Candidate, error) {
	r.log.InfoWithID(ctx, "[Repository: SearchCandidates] Called")
	var candidates []*models.Candidate
	if err := r.db.WithContext(ctx).Find(&candidates).Error; err != nil {
		r.log.ErrorWithID(ctx, "[Repository: SearchCandidates] Error Querying Candidates", err)
		return nil, err
	}
	return candidates, nil
}

func (r *CandidateRepository) GetCandidateByIds(ctx context.Context, candidateIDs []string) ([]*models.Candidate, error) {
	r.log.InfoWithID(ctx, "[Repository: GetCandidateByIds] Called")
	var candidates []*models.Candidate
	if err := r.db.WithContext(ctx).Where("id IN ?", candidateIDs).Find(&candidates).Error; err != nil {
		r.log.ErrorWithID(ctx, "[Repository: GetCandidateByIds] Error querying candidates", err)
		return nil, err
	}
	return candidates, nil
}
