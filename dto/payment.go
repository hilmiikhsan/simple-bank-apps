package dto

type PaymentRequest struct {
	AccountNumber string `json:"account_number" binding:"required,min=10,max=10"`
	Amount        int    `json:"amount" binding:"required"`
}
