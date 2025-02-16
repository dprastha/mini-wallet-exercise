package entities

import (
	"time"

	"github.com/google/uuid"
)

type WalletEntity struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CustomerID uuid.UUID `gorm:"type:uuid;not null" json:"customer_id"`
	Balance    float64   `gorm:"type:decimal(16,2);not null" json:"balance"`
	IsEnabled  bool      `gorm:"type:boolean;not null;default:false" json:"is_enabled"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (WalletEntity) TableName() string {
	return "wallets"
}
