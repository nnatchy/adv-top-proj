package models

import "time"

type User struct {
	ID            string    `gorm:"column:id;type:varchar(255);primaryKey"`
	WalletAddress string    `gorm:"column:wallet_address;type:varchar(255);not null;unique"`
	Username      string    `gorm:"column:username;type:varchar(100)"`
	CreatedAt     time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	TokenID       string    `gorm:"column:token_id;type:varchar(255);unique"`
	IsActive      bool      `gorm:"column:is_active;type:boolean;default:true"`
	StudentID     string    `gorm:"column:student_id;type:varchar(255)"`
	FirstName     string    `gorm:"column:first_name;type:varchar(255)"`
	LastName      string    `gorm:"column:last_name;type:varchar(255)"`
	HasVoteYear   bool      `gorm:"column:has_vote_year;type:boolean;default:false"`
	HasVoteRep    bool      `gorm:"column:has_vote_rep;type:boolean;default:false"`
}

func (User) TableName() string {
	return "users_adv"
}
