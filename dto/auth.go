package dto

import "time"

type Request struct {
	Username string `json:"username" binding:"required,min=5,max=20"`
	Password string `json:"password" binding:"required,min=8,max=100"`
}

type LoginResponse struct {
	ID       string    `json:"id"`
	Token    string    `json:"token"`
	Username string    `json:"username"`
	ExpireAt time.Time `json:"expire_at"`
}
