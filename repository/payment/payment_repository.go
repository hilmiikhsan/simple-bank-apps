package payment

import (
	"context"
	"database/sql"

	"github.com/simple-bank-apps/model"
)

type paymentRepository struct {
	db *sql.DB
}

func (p *paymentRepository) Create(ctx context.Context, payment model.Payment) error {
	_, err := p.db.ExecContext(ctx, createPaymentQuery, payment.CustomerID, payment.Amount)
	if err != nil {
		return err
	}

	return nil
}
