package auth

import (
	"context"

	"github.com/simple-bank-apps/config"
	"github.com/simple-bank-apps/dto"
	"github.com/simple-bank-apps/middleware"
	"github.com/simple-bank-apps/repository/customer"
)

type AuthUsecase interface {
	Register(ctx context.Context, req dto.RegisterRequest) error
	Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error)
	Logout(ctx context.Context, token *middleware.Token) error
}

func NewAuthUsecase(customerRepo customer.CustomerRepository, jwt middleware.JWT, cfg *config.Config) AuthUsecase {
	return &authUsecase{
		customerRepo: customerRepo,
		jwt:          jwt,
		cfg:          cfg,
	}
}
