package repo

import (
	"log"
	"mini-wallet-exercise/app/customer/domain"
	"mini-wallet-exercise/entities"
	"mini-wallet-exercise/interface/http/exception"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	model *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) domain.CustomerRepository {
	return &CustomerRepository{
		model: db.Model(&entities.CustomerEntity{}),
	}
}

func (r *CustomerRepository) Create(ctx *gin.Context, customer *entities.CustomerEntity) {
	err := r.model.WithContext(ctx).
		Create(customer).Error
	if err != nil {
		log.Println("error create customer: ", err)
		panic(*exception.ServerErrorException(err))
	}
}

func (r *CustomerRepository) FindOneByCustomerXID(ctx *gin.Context, customerXID uuid.UUID) *entities.CustomerEntity {
	var customer entities.CustomerEntity
	err := r.model.WithContext(ctx).
		Where("customer_xid = ?", customerXID).
		First(&customer).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	} else if err != nil {
		log.Println("error find one by customer xid: ", err)
		panic(*exception.ServerErrorException(err))
	}

	return &customer
}
