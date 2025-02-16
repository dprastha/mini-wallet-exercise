package entities

import (
	"time"

	"github.com/google/uuid"
)

type TransactionEntity struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	WalletID    uuid.UUID `gorm:"type:uuid;not null" json:"wallet_id"`
	Amount      float64   `gorm:"type:decimal(16,2);not null" json:"amount"`
	ReferenceID string    `gorm:"type:varchar(255);not null;uniqueIndex:idx_refid_type" json:"reference_id"`
	Type        string    `gorm:"type:varchar(255);not null;uniqueIndex:idx_refid_type" json:"type"`
	Status      string    `gorm:"type:varchar(255);not null" json:"status"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Relationship
	Wallet WalletEntity `gorm:"foreignKey:WalletID;references:ID" json:"wallet"`
}

func (TransactionEntity) TableName() string {
	return "transactions"
}
