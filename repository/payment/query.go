package payment

const (
	createPaymentQuery = `
	INSERT INTO payments (customer_id, amount) VALUES ($1, $2)
	`
)
