package main

import (
	"log"
	"mini-wallet-exercise/config"
	"mini-wallet-exercise/entities"
	"mini-wallet-exercise/interface/http/handler"
	dbconn "mini-wallet-exercise/pkg/db"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	customerRepo "mini-wallet-exercise/app/customer/repository"
	customerUC "mini-wallet-exercise/app/customer/usecase"
)

func initializeDatabase() *gorm.DB {
	db, err := dbconn.InitDb(&dbconn.DatabaseCredentials{
		Host:     config.DBHost,
		Username: config.DBUsername,
		Password: config.DBPassword,
		Port:     config.DBPort,
		Name:     config.DBName,
		TimeZone: config.DBTimeZone,
	})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func setupRouter() *gin.Engine {
	router := gin.New()

	router.Use(handler.RecoverPanic())

	return router
}

func initializeModule(db *gorm.DB, router *gin.Engine) {
	// Initialize repository
	customerRepository := customerRepo.NewCustomerRepository(db)

	// Initialize usecase
	customerUsecase := customerUC.NewCustomerUsecase(customerRepository)

	// Initialize handler
	handler.NewCustomerHandler(router, customerUsecase)
}

func main() {
	db := initializeDatabase()
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	db.AutoMigrate(&entities.CustomerEntity{}, &entities.WalletEntity{}, &entities.TransactionEntity{})

	router := setupRouter()

	initializeModule(db, router)

	if err := router.Run(":" + config.AppPort); err != nil {
		log.Fatalf("Failed to run server : %v", err)
	}

}
