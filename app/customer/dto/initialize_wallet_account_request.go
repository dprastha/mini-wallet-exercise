package dto

import "github.com/google/uuid"

type InitializeWalletAccountRequest struct {
	CustomerXID uuid.UUID `json:"customer_xid" binding:"required,uuid"`
}
