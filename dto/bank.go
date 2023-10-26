package dto

type BankResponse struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	AccountNumber string `json:"account_number"`
	AccountName   string `json:"account_name"`
}
