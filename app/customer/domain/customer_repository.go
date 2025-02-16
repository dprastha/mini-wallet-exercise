package domain

import (
	"mini-wallet-exercise/entities"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CustomerRepository interface {
	Create(ctx *gin.Context, customer *entities.CustomerEntity)
	FindOneByCustomerXID(ctx *gin.Context, customerXID uuid.UUID) *entities.CustomerEntity
}
