package bank

const (
	listBankQuery = `
	SELECT id, name, account_number, account_name FROM banks
	`

	getBankByAccountNumberQuery = `
	SELECT id, name, account_number, account_name FROM banks WHERE account_number = $1
	`
)
