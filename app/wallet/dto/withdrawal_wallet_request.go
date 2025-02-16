package dto

type WithdrawalWalletRequest struct {
	Amount      float64 `form:"amount" binding:"required,gt=0"`
	ReferenceID string  `form:"reference_id" binding:"required"`
}
