package auth

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/simple-bank-apps/config"
	"github.com/simple-bank-apps/dto"
	mockJWT "github.com/simple-bank-apps/middleware/mock"
	"github.com/simple-bank-apps/model"
	mockdatabase "github.com/simple-bank-apps/repository/customer/mock"
	"github.com/simple-bank-apps/utils"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCustomerRepo := mockdatabase.NewMockCustomerRepository(ctrl)

	authUC := NewAuthUsecase(mockCustomerRepo, nil, nil)

	mockCustomerRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return("", nil)
	mockCustomerRepo.EXPECT().List(gomock.Any()).Return(nil, nil)

	password, err := utils.HashPassword("12345678")
	assert.Nil(t, err)
	assert.NoError(t, err)

	request := dto.RegisterRequest{
		Username:      "User Test",
		Password:      password,
		AccountNumber: "1234567890",
		AccountName:   "Account Name",
	}

	err = authUC.Register(ctx, request)

	assert.Nil(t, err)
	assert.NoError(t, err)
}

func TestLogin(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCustomerRepo := mockdatabase.NewMockCustomerRepository(ctrl)
	mockJWT := mockJWT.NewMockJWT(ctrl)
	mockConfig := config.Config{}

	authUC := NewAuthUsecase(mockCustomerRepo, mockJWT, &mockConfig)

	password := "12345678"
	hashedPassword, err := utils.HashPassword(password)
	assert.NoError(t, err)

	requestRegister := dto.RegisterRequest{
		Username:      "User Test",
		Password:      password,
		AccountNumber: "1234567890",
		AccountName:   "Account Name",
	}

	mockCustomerRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return("", nil)
	mockCustomerRepo.EXPECT().List(gomock.Any()).Return(nil, nil)

	err = authUC.Register(ctx, requestRegister)
	assert.NoError(t, err)

	expectedCustomer := model.Customer{
		Username: "User Test",
		Password: hashedPassword,
	}

	mockCustomerRepo.EXPECT().GetByUsername(gomock.Any(), requestRegister.Username).Return(expectedCustomer, nil)

	requestLogin := dto.LoginRequest{
		Username: "User Test",
		Password: "12345678",
	}

	mockJWT.EXPECT().GetTokenFromRedis(gomock.Any(), gomock.Any(), gomock.Any()).Return("mockedToken", nil)

	response, err := authUC.Login(ctx, requestLogin)

	assert.Nil(t, err)
	assert.Equal(t, expectedCustomer.ID, response.ID)
}
