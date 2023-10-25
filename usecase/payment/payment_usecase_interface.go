package payment

import (
	"context"

	"github.com/simple-bank-apps/dto"
	"github.com/simple-bank-apps/repository/payment"
)

type PaymentUsecase interface {
	Payment(ctx context.Context, req dto.PaymentRequest, userID string) error
}

func NewPaymentUsecase(paymentRepo payment.PaymentRepository) PaymentUsecase {
	return &paymentUsecase{
		paymentRepo: paymentRepo,
	}
}
