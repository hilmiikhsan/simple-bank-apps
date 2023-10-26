package model

type Customer struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	Amount        int    `json:"amount"`
	AccountNumber string `json:"account_number"`
	AccountName   string `json:"account_name"`
}
