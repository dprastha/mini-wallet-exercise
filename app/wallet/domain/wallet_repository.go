package domain

import (
	"mini-wallet-exercise/entities"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type WalletRepository interface {
	FindOneByID(ctx *gin.Context, id uuid.UUID) *entities.WalletEntity
	FindOneByCustomerID(ctx *gin.Context, customerId uuid.UUID) *entities.WalletEntity
	Create(ctx *gin.Context, wallet *entities.WalletEntity)
	EnableByCustomerId(ctx *gin.Context, customerId uuid.UUID)
	DisableByCustomerId(ctx *gin.Context, customerId uuid.UUID)
	FindTransactionsByWalletID(ctx *gin.Context, walletID uuid.UUID) []entities.TransactionEntity
	CreateTransactionAndUpdateBalance(ctx *gin.Context, transaction *entities.TransactionEntity, wallet *entities.WalletEntity) *entities.TransactionEntity
	UpdateBalance(ctx *gin.Context, wallet *entities.WalletEntity)
}
