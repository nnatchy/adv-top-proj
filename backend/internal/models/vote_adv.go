package models

import (
	"time"
)

type Vote struct {
	ID            string    `gorm:"column:id;type:uuid;default:gen_random_uuid();primaryKey"` // Primary Key for vote
	StudentID     string    `gorm:"column:student_id;type:varchar(100);not null"`             // Foreign Key to User
	VoterAddress  string    `gorm:"column:voter_address;type:varchar(255);not null"`
	CandidateID   string    `gorm:"column:candidate_id;type:uuid;not null"`
	TransactionID string    `gorm:"column:transaction_id;type:varchar(255);not null;index"` // Ethereum transaction hash
	IsValid       bool      `gorm:"column:is_valid;type:boolean;default:false"`              // Validation status
	CreatedAt     time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
}

// TableName sets the custom table name for the Vote model
func (Vote) TableName() string {
	return "votes_adv"
}
