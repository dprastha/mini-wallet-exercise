package dto

type InitializeWalletAccountRequest struct {
	CustomerXID string `form:"customer_xid" binding:"required"`
}
