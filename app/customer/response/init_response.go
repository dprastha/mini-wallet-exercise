package response

import "github.com/google/uuid"

type InitResponse struct {
	Token uuid.UUID `json:"token"`
}
