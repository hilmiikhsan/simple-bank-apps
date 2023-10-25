package customer

const (
	createCustomerQuery = `
	INSERT INTO customers (username, password) VALUES ($1, $2)
	`

	listCustomerQuery = `
	SELECT id, username, password FROM customers
	`

	getCustomerByUsernameQuery = `
	SELECT id, username, password FROM customers WHERE username = $1
	`
)
