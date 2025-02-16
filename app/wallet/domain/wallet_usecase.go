package domain

import (
	"mini-wallet-exercise/entities"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type WalletUsecase interface {
	EnableByCustomerId(ctx *gin.Context, customerId uuid.UUID) error
	GetBalanceByCustomerId(ctx *gin.Context, customerId uuid.UUID) (float64, error)
	GetWalletTransactionsByCustomerId(ctx *gin.Context, customerId uuid.UUID) ([]entities.TransactionEntity, error)
	DepositWalletByCustomerId(ctx *gin.Context, customerId uuid.UUID, payload entities.TransactionEntity) (*entities.TransactionEntity, error)
	WithdrawWalletByCustomerId(ctx *gin.Context, customerId uuid.UUID, payload entities.TransactionEntity) (*entities.TransactionEntity, error)
	DisableWalletByCustomerId(ctx *gin.Context, customerId uuid.UUID) (*entities.WalletEntity, error)
}
