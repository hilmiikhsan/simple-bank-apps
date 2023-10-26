package customer

const (
	createCustomerQuery = `
	INSERT INTO customers (username, password, amount, account_number, account_name) VALUES ($1, $2, $3, $4, $5)
	RETURNING id
	`

	listCustomerQuery = `
	SELECT id, username, password, amount, account_number, account_name FROM customers
	`

	getCustomerByUsernameQuery = `
	SELECT id, username, password FROM customers WHERE username = $1
	`

	getCustomerByAccountNumberQuery = `
	SELECT id, username, amount, account_number, account_name FROM customers WHERE account_number = $1
	`

	getCustomerByIDQuery = `
	SELECT id, username, amount, account_number, account_name FROM customers WHERE id = $1
	`

	updateAmountByIDQuery = `
	UPDATE customers SET amount = $2 WHERE id = $1
	`
)
