package payment

import (
	"context"

	"github.com/simple-bank-apps/dto"
	"github.com/simple-bank-apps/repository/bank"
	"github.com/simple-bank-apps/repository/customer"
	"github.com/simple-bank-apps/repository/payment"
)

type PaymentUsecase interface {
	Payment(ctx context.Context, req dto.PaymentRequest) error
}

func NewPaymentUsecase(paymentRepo payment.PaymentRepository, customerRepo customer.CustomerRepository, bankRepo bank.BankRepository) PaymentUsecase {
	return &paymentUsecase{
		paymentRepo:  paymentRepo,
		customerRepo: customerRepo,
		bankRepo:     bankRepo,
	}
}
