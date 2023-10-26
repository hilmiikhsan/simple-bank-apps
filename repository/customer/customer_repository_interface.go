package customer

import (
	"context"
	"database/sql"

	"github.com/simple-bank-apps/model"
)

type CustomerRepository interface {
	Create(ctx context.Context, customer model.Customer) (string, error)
	List(ctx context.Context) ([]model.Customer, error)
	GetByUsername(ctx context.Context, username string) (model.Customer, error)
	GetByAccountNumber(ctx context.Context, accountNumber string) (model.Customer, error)
	GetByID(ctx context.Context, id string) (model.Customer, error)
	UpdateAmountByID(ctx context.Context, tx *sql.Tx, customer model.Customer) error
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRepository{
		db: db,
	}
}
