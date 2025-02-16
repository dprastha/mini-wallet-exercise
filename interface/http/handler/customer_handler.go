package handler

import (
	"mini-wallet-exercise/app/customer/domain"
	"mini-wallet-exercise/app/customer/dto"
	"mini-wallet-exercise/app/customer/response"
	"mini-wallet-exercise/entities"
	"mini-wallet-exercise/interface/http/middleware"
	"mini-wallet-exercise/interface/http/singleton"
	"mini-wallet-exercise/pkg"
	"mini-wallet-exercise/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		// Convert CustomerXID to UUID
		customerXID, err := uuid.Parse(request.CustomerXID)
		if err != nil {
			errorMessage := pkg.ErrorMessage{
				Error: "Invalid CustomerXID",
			}
			ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(errorMessage))
			return
		}

		customer := entities.CustomerEntity{
			CustomerXID: customerXID,
		}

		token := h.customerUsecase.Init(ctx, &customer)

		response := response.InitResponse{
			Token: token,
		}

		ctx.JSON(http.StatusOK, utils.SuccessResponse(response))
	}
}
