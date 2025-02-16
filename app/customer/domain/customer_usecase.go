package domain

import (
	"mini-wallet-exercise/entities"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CustomerUsecase interface {
	Init(ctx *gin.Context, customer *entities.CustomerEntity) uuid.UUID
}
