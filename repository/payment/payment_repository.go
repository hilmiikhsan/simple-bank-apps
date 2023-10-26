package payment

import (
	"context"
	"database/sql"

	"github.com/simple-bank-apps/model"
	"github.com/simple-bank-apps/repository/customer"
)

type paymentRepository struct {
	db           *sql.DB
	customerRepo customer.CustomerRepository
}

func (p *paymentRepository) Create(ctx context.Context, paymentModel model.Payment, customerModel model.Customer) error {
	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, createPaymentQuery, paymentModel.CustomerID, paymentModel.Amount, paymentModel.AccuntNumber)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = p.customerRepo.UpdateAmountByID(ctx, tx, customerModel)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
