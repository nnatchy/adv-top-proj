package models

import (
	"time"
)

type NFT struct {
	TokenID       string    `gorm:"column:token_id;type:varchar(255);primaryKey"`               // TokenID as the primary key
	WalletAddress string    `gorm:"column:wallet_address;type:varchar(255);not null"`           // Owner's wallet address
	NFTMetadata   string    `gorm:"column:nft_metadata;type:jsonb"`                             // Metadata stored in JSONB format
	CreatedAt     time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"` // Creation timestamp
}

func (NFT) TableName() string {
	return "nft_adv"
}
