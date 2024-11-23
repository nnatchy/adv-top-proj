package repositories

import (
	"context"
	"errors"

	"github.com/nnatchy/adv-top-proj/internal/infrastructure/log"
	"github.com/nnatchy/adv-top-proj/internal/models"
	"gorm.io/gorm"
)

type IVoteRepository interface {
	SaveVote(ctx context.Context, vote *models.Vote) error
	FindVoteByStudentID(ctx context.Context, studentID string) (*models.Vote, error)
	GetAllVotes(ctx context.Context) ([]models.Vote, error)
}

type VoteRepository struct {
	db  *gorm.DB
	log *log.Logger
}

func NewVoteRepository(db *gorm.DB, log *log.Logger) IVoteRepository {
	return &VoteRepository{
		db:  db,
		log: log,
	}
}

func (r *VoteRepository) SaveVote(ctx context.Context, vote *models.Vote) error {
	r.log.InfoWithID(ctx, "[Repository: SaveVote] Called")
	if err := r.db.Create(vote).Error; err != nil {
		r.log.ErrorWithID(ctx, "[Repository: SaveVote] Error saving vote", err)
		return err
	}
	r.log.InfoWithID(ctx, "[Repository: SaveVote] Vote saved successfully")
	return nil
}

func (r *VoteRepository) FindVoteByStudentID(ctx context.Context, studentID string) (*models.Vote, error) {
	r.log.InfoWithID(ctx, "[Repository: FindVoteByStudentID] Called")
	var vote models.Vote
	if err := r.db.Where("student_id = ?", studentID).First(&vote).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.log.ErrorWithID(ctx, "[Repository: FindVoteByStudentID] No vote found for StudentID")
			return nil, nil
		}
		r.log.ErrorWithID(ctx, "[Repository: FindVoteByStudentID] Error fetching vote", err)
		return nil, err
	}
	r.log.InfoWithID(ctx, "[Repository: FindVoteByStudentID] Vote found")
	return &vote, nil
}

// GetAllVotes retrieves all votes from the database
func (r *VoteRepository) GetAllVotes(ctx context.Context) ([]models.Vote, error) {
	r.log.InfoWithID(ctx, "[Repository: GetAllVotes] Fetching all votes")

	var votes []models.Vote
	if err := r.db.WithContext(ctx).Session(&gorm.Session{PrepareStmt: true}).Find(&votes).Error; err != nil {
		r.log.ErrorWithID(ctx, "[Repository: GetAllVotes] Failed to fetch votes", err)
		return nil, err
	}

	return votes, nil
}

// hasvoted
