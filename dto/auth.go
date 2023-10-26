package dto

import "time"

type RegisterRequest struct {
	Username      string `json:"username" binding:"required,min=5,max=20"`
	Password      string `json:"password" binding:"required,min=8,max=100"`
	AccountNumber string `json:"account_number" binding:"required,min=10,max=10"`
	AccountName   string `json:"account_name" binding:"required,min=1,max=20"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required,min=5,max=20"`
	Password string `json:"password" binding:"required,min=8,max=100"`
}

type LoginResponse struct {
	ID       string    `json:"id"`
	Token    string    `json:"token"`
	Username string    `json:"username"`
	ExpireAt time.Time `json:"expire_at"`
}
