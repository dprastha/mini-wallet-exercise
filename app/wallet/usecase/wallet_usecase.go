package usecase

import (
	"errors"
	"mini-wallet-exercise/app/wallet/domain"
	"mini-wallet-exercise/entities"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type WalletUsecase struct {
	repo domain.WalletRepository
}

func NewWalletUsecase(walletRepo domain.WalletRepository) domain.WalletUsecase {
	return &WalletUsecase{
		repo: walletRepo,
	}
}

func (uc *WalletUsecase) EnableByCustomerId(ctx *gin.Context, customerId uuid.UUID) error {
	// Find wallet by customerID
	existingWallet := uc.repo.FindOneByCustomerID(ctx, customerId)
	if existingWallet == nil {
		// Create new wallet
		newWallet := entities.WalletEntity{
			CustomerID: customerId,
			IsEnabled:  true,
			Balance:    0,
		}
		uc.repo.Create(ctx, &newWallet)
	} else if !existingWallet.IsEnabled {
		// Enable existing wallet
		uc.repo.EnableByCustomerId(ctx, customerId)
	} else if existingWallet.IsEnabled {
		// Do nothing
		return errors.New("already enabled")
	}

	return nil
}

func (uc *WalletUsecase) GetBalanceByCustomerId(ctx *gin.Context, customerId uuid.UUID) (float64, error) {
	// Find wallet by customerID
	existingWallet := uc.repo.FindOneByCustomerID(ctx, customerId)
	if existingWallet == nil {
		return 0, errors.New("wallet not found")
	} else if !existingWallet.IsEnabled {
		return 0, errors.New("wallet is disabled")
	}

	return existingWallet.Balance, nil
}

func (uc *WalletUsecase) GetWalletTransactionsByCustomerId(ctx *gin.Context, customerId uuid.UUID) ([]entities.TransactionEntity, error) {
	// Find wallet by customerID
	existingWallet := uc.repo.FindOneByCustomerID(ctx, customerId)
	if existingWallet == nil {
		return nil, errors.New("wallet not found")
	} else if !existingWallet.IsEnabled {
		return nil, errors.New("wallet is disabled")
	}

	// Find transactions by walletID
	transactions := uc.repo.FindTransactionsByWalletID(ctx, existingWallet.ID)

	return transactions, nil
}

func (uc *WalletUsecase) DepositWalletByCustomerId(ctx *gin.Context, customerId uuid.UUID, payload entities.TransactionEntity) (*entities.TransactionEntity, error) {
	// Find wallet by customerID
	existingWallet := uc.repo.FindOneByCustomerID(ctx, customerId)
	if existingWallet == nil {
		return nil, errors.New("wallet not found")
	} else if !existingWallet.IsEnabled {
		return nil, errors.New("wallet is disabled")
	}

	// Assign payload
	payload.WalletID = existingWallet.ID

	// Create transaction and update balance
	transaction := uc.repo.CreateTransactionAndUpdateBalance(ctx, &payload, existingWallet)

	return transaction, nil
}

func (uc *WalletUsecase) WithdrawWalletByCustomerId(ctx *gin.Context, customerId uuid.UUID, payload entities.TransactionEntity) (*entities.TransactionEntity, error) {
	// Find wallet by customerID
	existingWallet := uc.repo.FindOneByCustomerID(ctx, customerId)
	if existingWallet == nil {
		return nil, errors.New("wallet not found")
	} else if !existingWallet.IsEnabled {
		return nil, errors.New("wallet is disabled")
	}

	// Assign payload
	payload.WalletID = existingWallet.ID

	// Check if balance is sufficient
	if existingWallet.Balance < payload.Amount {
		return nil, errors.New("insufficient balance")
	}

	// Create transaction and update balance
	transaction := uc.repo.CreateTransactionAndUpdateBalance(ctx, &payload, existingWallet)

	return transaction, nil
}

func (uc *WalletUsecase) DisableWalletByCustomerId(ctx *gin.Context, customerId uuid.UUID) (*entities.WalletEntity, error) {
	// Find wallet by customerID
	existingWallet := uc.repo.FindOneByCustomerID(ctx, customerId)
	if existingWallet == nil {
		return nil, errors.New("wallet not found")
	} else if !existingWallet.IsEnabled {
		return nil, errors.New("wallet is already disabled")
	}

	// Disable wallet
	uc.repo.DisableByCustomerId(ctx, customerId)

	disabledWallet := uc.repo.FindOneByCustomerID(ctx, customerId)

	return disabledWallet, nil
}
