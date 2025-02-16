package handler

import (
	"mini-wallet-exercise/app/wallet/constant"
	"mini-wallet-exercise/app/wallet/domain"
	"mini-wallet-exercise/app/wallet/dto"
	"mini-wallet-exercise/app/wallet/response"
	"mini-wallet-exercise/entities"
	"mini-wallet-exercise/interface/http/guard"
	"mini-wallet-exercise/interface/http/middleware"
	"mini-wallet-exercise/interface/http/singleton"
	"mini-wallet-exercise/pkg"
	"mini-wallet-exercise/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WalletHandler struct {
	walletUsecase domain.WalletUsecase
}

func NewWalletHandler(router *gin.Engine, walletUsecase domain.WalletUsecase) {
	walletHandlerRoute := router.Group("/api/v1/wallet", guard.AuthGuard())

	walletHandler := &WalletHandler{
		walletUsecase: walletUsecase,
	}

	walletHandlerRoute.POST("", walletHandler.Enable())
	walletHandlerRoute.GET("", walletHandler.GetBalance())
	walletHandlerRoute.GET("/transactions", walletHandler.GetWalletTransactions())
	walletHandlerRoute.POST("/deposit", middleware.ValidateRequestFormData[dto.DepositWalletRequest](), walletHandler.Deposit())
	walletHandlerRoute.POST("/withdrawal", middleware.ValidateRequestFormData[dto.WithdrawalWalletRequest](), walletHandler.Withdrawal())
	walletHandlerRoute.PATCH("", walletHandler.Disable())
}

func (h *WalletHandler) Enable() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		customer := ctx.MustGet("customer").(entities.CustomerEntity)

		err := h.walletUsecase.EnableByCustomerId(ctx, customer.ID)

		if err != nil {
			errorMessage := pkg.ErrorMessage{
				Error: err.Error(),
			}
			ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(errorMessage))
			return
		}

		ctx.JSON(http.StatusOK, utils.SuccessResponse("Wallet has been enabled"))
	}
}

func (h *WalletHandler) GetBalance() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		customer := ctx.MustGet("customer").(entities.CustomerEntity)

		balance, err := h.walletUsecase.GetBalanceByCustomerId(ctx, customer.ID)

		if err != nil {
			errorMessage := pkg.ErrorMessage{
				Error: err.Error(),
			}
			ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(errorMessage))
			return
		}

		ctx.JSON(http.StatusOK, utils.SuccessResponse(balance))
	}
}

func (h *WalletHandler) GetWalletTransactions() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		customer := ctx.MustGet("customer").(entities.CustomerEntity)

		transactions, err := h.walletUsecase.GetWalletTransactionsByCustomerId(ctx, customer.ID)

		if err != nil {
			errorMessage := pkg.ErrorMessage{
				Error: err.Error(),
			}
			ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(errorMessage))
			return
		}

		ctx.JSON(http.StatusOK, utils.SuccessResponse(transactions))
	}
}

func (h *WalletHandler) Deposit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		customer := ctx.MustGet("customer").(entities.CustomerEntity)
		request := singleton.GetHTTPRequest[dto.DepositWalletRequest](ctx)

		payload := entities.TransactionEntity{
			Amount:      request.Amount,
			Type:        constant.TransactionTypeDeposit,
			Status:      constant.TransactionStatusSuccess,
			ReferenceID: request.ReferenceID,
		}

		transaction, err := h.walletUsecase.DepositWalletByCustomerId(ctx, customer.ID, payload)

		if err != nil {
			errorMessage := pkg.ErrorMessage{
				Error: err.Error(),
			}
			ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(errorMessage))
			return
		}

		depositResponse := response.DepositResponse{
			ID:          transaction.ID,
			DepositedBy: customer.ID,
			Status:      transaction.Status,
			DepositedAt: transaction.CreatedAt,
			Amount:      transaction.Amount,
			ReferenceID: transaction.ReferenceID,
		}

		ctx.JSON(http.StatusOK, utils.SuccessResponse(depositResponse))
	}
}

func (h *WalletHandler) Withdrawal() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		customer := ctx.MustGet("customer").(entities.CustomerEntity)
		request := singleton.GetHTTPRequest[dto.WithdrawalWalletRequest](ctx)

		payload := entities.TransactionEntity{
			Amount:      request.Amount,
			Type:        constant.TransactionTypeWithdraw,
			Status:      constant.TransactionStatusSuccess,
			ReferenceID: request.ReferenceID,
		}

		transaction, err := h.walletUsecase.WithdrawWalletByCustomerId(ctx, customer.ID, payload)

		if err != nil {
			errorMessage := pkg.ErrorMessage{
				Error: err.Error(),
			}
			ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(errorMessage))
			return
		}

		withdrawalResponse := response.WithdrawalResponse{
			ID:          transaction.ID,
			WithdrawnBy: customer.ID,
			Status:      transaction.Status,
			WithdrawnAt: transaction.CreatedAt,
			Amount:      transaction.Amount,
			ReferenceID: transaction.ReferenceID,
		}

		ctx.JSON(http.StatusOK, utils.SuccessResponse(withdrawalResponse))
	}
}

func (h *WalletHandler) Disable() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		customer := ctx.MustGet("customer").(entities.CustomerEntity)

		disabledWallet, err := h.walletUsecase.DisableWalletByCustomerId(ctx, customer.ID)

		if err != nil {
			errorMessage := pkg.ErrorMessage{
				Error: err.Error(),
			}
			ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(errorMessage))
			return
		}

		ctx.JSON(http.StatusOK, utils.SuccessResponse(disabledWallet))
	}
}
