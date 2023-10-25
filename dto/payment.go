package dto

type PaymentRequest struct {
	Amount int `json:"amount" binding:"required"`
}
