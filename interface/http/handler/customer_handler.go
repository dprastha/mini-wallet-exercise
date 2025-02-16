package handler

import (
	"mini-wallet-exercise/app/customer/domain"
	"mini-wallet-exercise/app/customer/dto"
	"mini-wallet-exercise/entities"
	"mini-wallet-exercise/interface/http/middleware"
	"mini-wallet-exercise/interface/http/singleton"
	"mini-wallet-exercise/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	customerUsecase domain.CustomerUsecase
}

func NewCustomerHandler(router *gin.Engine, customerUsecase domain.CustomerUsecase) {
	customerHandlerRoute := router.Group("/api/v1/init")

	customerHandler := &CustomerHandler{
		customerUsecase: customerUsecase,
	}

	customerHandlerRoute.POST("", middleware.ValidateRequestFormData[dto.InitializeWalletAccountRequest](), customerHandler.Init())
}

func (h *CustomerHandler) Init() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := singleton.GetHTTPRequest[dto.InitializeWalletAccountRequest](ctx)

		customer := entities.CustomerEntity{
			CustomerXID: request.CustomerXID,
		}

		token := h.customerUsecase.Init(ctx, &customer)

		ctx.JSON(http.StatusOK, utils.SuccessResponse(token))
	}
}
