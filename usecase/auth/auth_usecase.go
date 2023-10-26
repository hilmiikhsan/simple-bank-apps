package auth

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/simple-bank-apps/config"
	"github.com/simple-bank-apps/constants"
	"github.com/simple-bank-apps/dto"
	"github.com/simple-bank-apps/middleware"
	"github.com/simple-bank-apps/model"
	"github.com/simple-bank-apps/repository/customer"
	"github.com/simple-bank-apps/utils"
)

type authUsecase struct {
	customerRepo customer.CustomerRepository
	jwt          middleware.JWT
	cfg          *config.Config
}

func (a *authUsecase) Register(ctx context.Context, req dto.RegisterRequest) error {
	password, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}

	customers, err := a.customerRepo.List(ctx)
	if err != nil {
		return err
	}

	for _, customer := range customers {
		if customer.Username == req.Username {
			return constants.ErrUsernameAlreadyExist
		}

		if customer.AccountNumber == req.AccountNumber {
			return constants.ErrAccountNumberAlreadyExist
		}
	}

	_, err = a.customerRepo.Create(ctx, model.Customer{
		Username:      req.Username,
		Password:      password,
		Amount:        50000,
		AccountNumber: req.AccountNumber,
		AccountName:   req.AccountName,
	})
	if err != nil {
		return err
	}

	return nil
}

func (a *authUsecase) Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error) {
	response := dto.LoginResponse{}
	var token string

	customer, err := a.customerRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response, constants.ErrUsernameOrPasswordNotMatch
		}

		return response, err
	}

	if !utils.CheckPasswordHash(req.Password, customer.Password) {
		return response, constants.ErrUsernameOrPasswordNotMatch
	}

	token, err = a.jwt.GetTokenFromRedis(ctx, customer.ID, "auth-customer")
	if err != nil {
		return response, err
	}

	if token == "" {
		token, err = a.jwt.GenerateJWTToken(ctx, "auth-customer", middleware.JWTRequest{
			ID:       customer.ID,
			Username: customer.Username,
		})
		if err != nil {
			return response, err
		}
	}

	response = dto.LoginResponse{
		ID:       customer.ID,
		Token:    token,
		Username: customer.Username,
		ExpireAt: time.Now().Add(time.Duration(a.cfg.JWT.TokenLifeTimeHour) * time.Hour),
	}

	return response, nil
}

func (a *authUsecase) Logout(ctx context.Context, token *middleware.Token) error {
	claims, err := a.jwt.ExtractJWTClaims(ctx, token.Value)
	if err != nil {
		return err
	}

	err = a.jwt.DeleteTokenFromRedis(ctx, claims.ID, "auth-customer")
	if err != nil {
		return err
	}

	return nil
}
