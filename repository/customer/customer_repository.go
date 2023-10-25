package customer

import (
	"context"
	"database/sql"

	"github.com/simple-bank-apps/model"
)

type customerRepository struct {
	db *sql.DB
}

func (c *customerRepository) Create(ctx context.Context, customer model.Customer) error {
	_, err := c.db.ExecContext(ctx, createCustomerQuery, customer.Username, customer.Password)
	if err != nil {
		return err
	}

	return nil
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
