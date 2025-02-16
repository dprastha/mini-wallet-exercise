package dto

type DisableWalletRequest struct {
	IsDisabled bool `form:"is_disabled" binding:"required"`
}
