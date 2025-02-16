package usecase

import (
	"mini-wallet-exercise/app/customer/domain"
	"mini-wallet-exercise/entities"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CustomerUsecase struct {
	repo domain.CustomerRepository
}

func NewCustomerUsecase(repo domain.CustomerRepository) domain.CustomerUsecase {
	return &CustomerUsecase{
		repo: repo,
	}
}

func (uc *CustomerUsecase) Init(ctx *gin.Context, customer *entities.CustomerEntity) uuid.UUID {
	// Find customer by customerXID
	existingCustomer := uc.repo.FindOneByCustomerXID(ctx, customer.CustomerXID)

	// If customer already exists, return the token
	if existingCustomer != nil {
		return existingCustomer.Token
	}

	// Create new customer
	token := uuid.New()
	customer.Token = token
	uc.repo.Create(ctx, customer)

	return token
}
