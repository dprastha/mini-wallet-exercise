package entities

import (
	"time"

	"github.com/google/uuid"
)

type CustomerEntity struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CustomerXID uuid.UUID `gorm:"column:customer_xid;type:uuid;not null;unique" json:"customer_xid"`
	Token       uuid.UUID `gorm:"type:uuid;not null;unique" json:"token"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (CustomerEntity) TableName() string {
	return "customers"
}
