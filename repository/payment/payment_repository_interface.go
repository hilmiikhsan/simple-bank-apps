package payment

import (
	"context"
	"database/sql"

	"github.com/simple-bank-apps/model"
)

type PaymentRepository interface {
	Create(ctx context.Context, payment model.Payment) error
}

func NewPaymentRepository(db *sql.DB) PaymentRepository {
	return &paymentRepository{
		db: db,
	}
}
