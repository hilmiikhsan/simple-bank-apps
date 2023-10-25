package payment

import (
	"context"

	"github.com/simple-bank-apps/dto"
	"github.com/simple-bank-apps/model"
	"github.com/simple-bank-apps/repository/payment"
)

type paymentUsecase struct {
	paymentRepo payment.PaymentRepository
}

func (p *paymentUsecase) Payment(ctx context.Context, req dto.PaymentRequest, userID string) error {
	err := p.paymentRepo.Create(ctx, model.Payment{
		CustomerID: userID,
		Amount:     req.Amount,
	})
	if err != nil {
		return err
	}

	return nil
}
