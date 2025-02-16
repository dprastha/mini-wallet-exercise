package response

import (
	"time"

	"github.com/google/uuid"
)

type DepositResponse struct {
	ID          uuid.UUID `json:"id"`
	DepositedBy uuid.UUID `json:"deposited_by"`
	Status      string    `json:"status"`
	DepositedAt time.Time `json:"deposited_at"`
	Amount      float64   `json:"amount"`
	ReferenceID string    `json:"reference_id"`
}
