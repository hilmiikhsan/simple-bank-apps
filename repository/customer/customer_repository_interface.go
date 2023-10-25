package customer

import (
	"context"
	"database/sql"

	"github.com/simple-bank-apps/model"
)

type CustomerRepository interface {
	Create(ctx context.Context, customer model.Customer) error
	List(ctx context.Context) ([]model.Customer, error)
	GetByUsername(ctx context.Context, username string) (model.Customer, error)
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRepository{
		db: db,
	}
}
