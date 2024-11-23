package models

import (
	"time"

	"github.com/google/uuid"
)

type Candidate struct {
	ID        uuid.UUID `gorm:"type:uuid;column:id;default:gen_random_uuid();primaryKey"`
	FirstName string    `gorm:"column:first_name;type:varchar(255);not null"`
	LastName  string    `gorm:"column:last_name;type:varchar(255);not null"`
	StudentID string    `gorm:"column:student_id;type:varchar(100); not null"`
	Position  string    `gorm:"column:position;type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	Major     string    `gorm:"column:major;type:varchar(100);not null"`
}

func (Candidate) TableName() string {
	return "candidate_adv"
}
