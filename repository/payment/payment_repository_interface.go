package payment

import (
	"context"
	"database/sql"

	"github.com/simple-bank-apps/model"
	"github.com/simple-bank-apps/repository/customer"
)

type PaymentRepository interface {
	Create(ctx context.Context, paymentModel model.Payment, customerModel model.Customer) error
}

func NewPaymentRepository(db *sql.DB, customerRepo customer.CustomerRepository) PaymentRepository {
	return &paymentRepository{
		db:           db,
		customerRepo: customerRepo,
	}
}
