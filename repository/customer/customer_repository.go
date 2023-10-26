package customer

import (
	"context"
	"database/sql"
	"errors"

	"github.com/simple-bank-apps/model"
)

type customerRepository struct {
	db *sql.DB
}

func (c *customerRepository) Create(ctx context.Context, customer model.Customer) (string, error) {
	var id string
	err := c.db.QueryRowContext(
		ctx,
		createCustomerQuery,
		customer.Username,
		customer.Password,
		customer.Amount,
		customer.AccountNumber,
		customer.AccountName,
	).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (c *customerRepository) List(ctx context.Context) ([]model.Customer, error) {
	rows, err := c.db.QueryContext(ctx, listCustomerQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []model.Customer
	for rows.Next() {
		var customer model.Customer
		err := rows.Scan(
			&customer.ID,
			&customer.Username,
			&customer.Password,
			&customer.Amount,
			&customer.AccountNumber,
			&customer.AccountName,
		)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

func (c *customerRepository) GetByUsername(ctx context.Context, username string) (model.Customer, error) {
	var customer model.Customer
	err := c.db.QueryRowContext(ctx, getCustomerByUsernameQuery, username).Scan(
		&customer.ID,
		&customer.Username,
		&customer.Password,
	)
	if err != nil {
		return model.Customer{}, err
	}

	return customer, nil
}

func (c *customerRepository) GetByAccountNumber(ctx context.Context, accountNumber string) (model.Customer, error) {
	var customer model.Customer
	err := c.db.QueryRowContext(ctx, getCustomerByAccountNumberQuery, accountNumber).Scan(
		&customer.ID,
		&customer.Username,
		&customer.Amount,
		&customer.AccountNumber,
		&customer.AccountName,
	)
	if err != nil {
		return model.Customer{}, err
	}

	return customer, nil
}

func (c *customerRepository) GetByID(ctx context.Context, id string) (model.Customer, error) {
	var customer model.Customer
	err := c.db.QueryRowContext(ctx, getCustomerByIDQuery, id).Scan(
		&customer.ID,
		&customer.Username,
		&customer.Amount,
		&customer.AccountNumber,
		&customer.AccountName,
	)
	if err != nil {
		return model.Customer{}, err
	}

	return customer, nil
}

func (c *customerRepository) UpdateAmountByID(ctx context.Context, tx *sql.Tx, customer model.Customer) error {
	result, err := tx.ExecContext(ctx, updateAmountByIDQuery, customer.ID, customer.Amount)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows updated")
	}

	return nil
}
