package payment

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/simple-bank-apps/config"
	"github.com/simple-bank-apps/dto"
	mockJWT "github.com/simple-bank-apps/middleware/mock"
	"github.com/simple-bank-apps/model"
	mockCustomer "github.com/simple-bank-apps/repository/customer/mock"
	mockPayment "github.com/simple-bank-apps/repository/payment/mock"
	"github.com/simple-bank-apps/usecase/auth"
	"github.com/simple-bank-apps/utils"
	"github.com/stretchr/testify/assert"
)

func TestPayement(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPaymentRepo := mockPayment.NewMockPaymentRepository(ctrl)
	mockCustomerRepo := mockCustomer.NewMockCustomerRepository(ctrl)
	mockJWT := mockJWT.NewMockJWT(ctrl)
	mockConfig := config.Config{}

	authUC := auth.NewAuthUsecase(mockCustomerRepo, mockJWT, &mockConfig)

	password := "12345678"
	hashedPassword, err := utils.HashPassword(password)
	assert.NoError(t, err)

	requestRegister := dto.Request{
		Username: "User Test",
		Password: password,
	}

	mockCustomerRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
	mockCustomerRepo.EXPECT().List(gomock.Any()).Return(nil, nil)

	err = authUC.Register(ctx, requestRegister)
	assert.Nil(t, err)
	assert.NoError(t, err)

	expectedCustomer := model.Customer{
		Username: "User Test",
		Password: hashedPassword,
	}

	mockCustomerRepo.EXPECT().GetByUsername(gomock.Any(), requestRegister.Username).Return(expectedCustomer, nil)

	requestLogin := dto.Request{
		Username: "User Test",
		Password: "12345678",
	}

	mockJWT.EXPECT().GetTokenFromRedis(gomock.Any(), gomock.Any(), gomock.Any()).Return("mockedToken", nil)

	response, err := authUC.Login(ctx, requestLogin)

	assert.Nil(t, err)
	assert.Equal(t, expectedCustomer.ID, response.ID)

	paymentUC := NewPaymentUsecase(mockPaymentRepo)

	mockPaymentRepo.EXPECT().Create(ctx, gomock.Any()).Return(nil)

	err = paymentUC.Payment(ctx, dto.PaymentRequest{
		Amount: 1000,
	}, response.ID)

	assert.Nil(t, err)
	assert.NoError(t, err)
}
