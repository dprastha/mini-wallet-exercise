package repository

import (
	"log"
	"mini-wallet-exercise/app/wallet/domain"
	"mini-wallet-exercise/entities"
	"mini-wallet-exercise/interface/http/exception"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WalletRepository struct {
	model            *gorm.DB
	transactionModel *gorm.DB
}

func NewWalletRepository(db *gorm.DB) domain.WalletRepository {
	return &WalletRepository{
		model:            db,
		transactionModel: db,
	}
}

func (r *WalletRepository) FindOneByID(ctx *gin.Context, id uuid.UUID) *entities.WalletEntity {
	var wallet entities.WalletEntity
	err := r.model.WithContext(ctx).
		First(&wallet, id).Error

	if err == gorm.ErrRecordNotFound {
		return nil
	} else if err != nil {
		log.Println("error find one by id: ", err)
		panic(*exception.ServerErrorException(err))
	}

	return &wallet
}

func (r *WalletRepository) FindOneByCustomerID(ctx *gin.Context, customerId uuid.UUID) *entities.WalletEntity {
	var wallet entities.WalletEntity
	err := r.model.WithContext(ctx).
		Where("customer_id = ?", customerId).
		First(&wallet).Error

	if err == gorm.ErrRecordNotFound {
		return nil
	} else if err != nil {
		log.Println("error find one by customer id: ", err)
		panic(*exception.ServerErrorException(err))
	}

	return &wallet
}

func (r *WalletRepository) Create(ctx *gin.Context, wallet *entities.WalletEntity) {
	err := r.model.WithContext(ctx).
		Create(wallet).Error

	if err != nil {
		log.Println("error create wallet: ", err)
		panic(*exception.ServerErrorException(err))
	}
}

func (r *WalletRepository) EnableByCustomerId(ctx *gin.Context, customerId uuid.UUID) {
	err := r.model.WithContext(ctx).
		Where("customer_id = ?", customerId).
		Updates(map[string]interface{}{
			"is_active":   false,
			"disabled_at": nil,
		}).Error

	if err != nil {
		log.Println("error enable wallet: ", err)
		panic(*exception.ServerErrorException(err))
	}
}

func (r *WalletRepository) DisableByCustomerId(ctx *gin.Context, customerId uuid.UUID) {
	err := r.model.WithContext(ctx).
		Where("customer_id = ?", customerId).
		Updates(map[string]interface{}{
			"is_active":   false,
			"disabled_at": time.Now(),
		}).Error

	if err != nil {
		log.Println("error disable wallet: ", err)
		panic(*exception.ServerErrorException(err))
	}
}

func (r *WalletRepository) FindTransactionsByWalletID(ctx *gin.Context, walletID uuid.UUID) []entities.TransactionEntity {
	var transactions []entities.TransactionEntity
	err := r.transactionModel.WithContext(ctx).
		Where("wallet_id = ?", walletID).
		Find(&transactions).Error

	if err != nil {
		log.Println("error find transactions by wallet id: ", err)
		panic(*exception.ServerErrorException(err))
	}

	return transactions
}

func (r *WalletRepository) CreateTransactionAndUpdateBalance(ctx *gin.Context, transaction *entities.TransactionEntity, wallet *entities.WalletEntity) *entities.TransactionEntity {
	// Use database transaction to ensure data consistency
	tx := r.model.WithContext(ctx).Begin()

	err := tx.Create(transaction).Error
	if err != nil {
		tx.Rollback()
		log.Println("error create transaction: ", err)
		panic(*exception.ServerErrorException(err))
	}

	err = tx.Model(wallet).
		Updates(entities.WalletEntity{Balance: wallet.Balance}).Error
	if err != nil {
		tx.Rollback()
		log.Println("error update balance: ", err)
		panic(*exception.ServerErrorException(err))
	}

	tx.Commit()
	return transaction
}

func (r *WalletRepository) UpdateBalance(ctx *gin.Context, wallet *entities.WalletEntity) {
	err := r.model.WithContext(ctx).
		Model(wallet).
		Updates(entities.WalletEntity{Balance: wallet.Balance}).Error

	if err != nil {
		log.Println("error update balance: ", err)
		panic(*exception.ServerErrorException(err))
	}
}
