package payment

import (
	"context"
	"database/sql"
	"errors"

	"github.com/simple-bank-apps/constants"
	"github.com/simple-bank-apps/dto"
	"github.com/simple-bank-apps/middleware"
	"github.com/simple-bank-apps/model"
	"github.com/simple-bank-apps/repository/bank"
	"github.com/simple-bank-apps/repository/customer"
	"github.com/simple-bank-apps/repository/payment"
)

type paymentUsecase struct {
	paymentRepo  payment.PaymentRepository
	customerRepo customer.CustomerRepository
	bankRepo     bank.BankRepository
}

func (p *paymentUsecase) Payment(ctx context.Context, req dto.PaymentRequest) error {
	bank, err := p.bankRepo.GetByAccountNumber(ctx, req.AccountNumber)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return constants.ErrAccountNumberNotFound
		}
		return err
	}

	customer, err := p.customerRepo.GetByID(ctx, middleware.GetTokenMiddlewareFromContext(ctx).ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return constants.ErrCustomerNotFound
		}
		return err
	}

	if req.Amount > customer.Amount {
		return constants.ErrAmountNotEnough
	}

	paymentModel := model.Payment{
		CustomerID:   customer.ID,
		Amount:       req.Amount,
		AccuntNumber: bank.AccountNumber,
	}

	customerModel := model.Customer{
		ID:     customer.ID,
		Amount: customer.Amount - req.Amount,
	}

	err = p.paymentRepo.Create(ctx, paymentModel, customerModel)
	if err != nil {
		return err
	}

	return nil
}
