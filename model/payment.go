package model

type Payment struct {
	ID           string `json:"id"`
	CustomerID   string `json:"customer_id"`
	Amount       int    `json:"amount"`
	AccuntNumber string `json:"account_number"`
}
