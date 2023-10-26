package bank

import (
	"context"
	"database/sql"

	"github.com/simple-bank-apps/model"
)

type BankRepository interface {
	List(ctx context.Context) ([]model.Bank, error)
	GetByAccountNumber(ctx context.Context, accountNumber string) (model.Bank, error)
}

func NewBankRepository(db *sql.DB) BankRepository {
	return &bankRepository{
		db: db,
	}
}
