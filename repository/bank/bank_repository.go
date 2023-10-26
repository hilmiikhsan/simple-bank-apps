package bank

import (
	"context"
	"database/sql"

	"github.com/simple-bank-apps/model"
)

type bankRepository struct {
	db *sql.DB
}

func (b *bankRepository) List(ctx context.Context) ([]model.Bank, error) {
	rows, err := b.db.QueryContext(ctx, listBankQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var banks []model.Bank
	for rows.Next() {
		var bank model.Bank
		err := rows.Scan(
			&bank.ID,
			&bank.Name,
			&bank.AccountNumber,
			&bank.AccountName,
		)
		if err != nil {
			return nil, err
		}
		banks = append(banks, bank)
	}

	return banks, nil
}

func (b *bankRepository) GetByAccountNumber(ctx context.Context, accountNumber string) (model.Bank, error) {
	var bank model.Bank
	err := b.db.QueryRowContext(ctx, getBankByAccountNumberQuery, accountNumber).Scan(
		&bank.ID,
		&bank.Name,
		&bank.AccountNumber,
		&bank.AccountName,
	)
	if err != nil {
		return model.Bank{}, err
	}

	return bank, nil
}
