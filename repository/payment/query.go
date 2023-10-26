package payment

const (
	createPaymentQuery = `
	INSERT INTO payments (customer_id, amount, account_number) VALUES ($1, $2, $3)`
)
