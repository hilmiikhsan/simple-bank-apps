package customer

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/simple-bank-apps/model"
	"github.com/simple-bank-apps/utils"
	"github.com/stretchr/testify/require"
)

var TestDb *sql.DB

func SetupTestDb() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost", "5432", "postgres", "21012123op", "bank_apps_test",
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("failed to connect to database")
		os.Exit(1)
	}

	if err = db.Ping(); err != nil {
		fmt.Println("failed to ping database")
		os.Exit(1)
	}
	TestDb = db
}

func TestMain(m *testing.M) {
	SetupTestDb()
	code := m.Run()
	os.Exit(code)
}

// func createRandomCustomer(t *testing.T, customerRepo CustomerRepository) string {
// 	ctx := context.Background()
// 	password, err := utils.HashPassword("12345678")
// 	require.NoError(t, err)

// 	arg := model.Customer{
// 		Username:      "User Test",
// 		Password:      password,
// 		Amount:        50000,
// 		AccountNumber: "1234567890",
// 		AccountName:   "Account Name",
// 	}

// 	id, err := customerRepo.Create(ctx, arg)

// 	require.NoError(t, err)
// 	require.Equal(t, arg.Username, "User Test")
// 	require.Equal(t, arg.Password, password)

// 	return id
// }

func clearData(t *testing.T, db *sql.DB) {
	_, err := db.Exec("DELETE FROM customers WHERE username = 'User Test'")
	if err != nil {
		t.Fatalf("failed to clear data: %v", err)
	}
}

func TestCustomerRepository_Success(t *testing.T) {
	if testing.Short() {
		os.Exit(0)
	}

	ctx := context.Background()
	customerRepo := NewCustomerRepository(TestDb)

	password, err := utils.HashPassword("12345678")
	require.NoError(t, err)

	arg := model.Customer{
		Username:      "User Test 2",
		Password:      password,
		Amount:        50000,
		AccountNumber: "1234567891",
		AccountName:   "Account Name",
	}

	id, _ := customerRepo.Create(ctx, arg)

	t.Run("Create customer", func(t *testing.T) {
		arg2 := model.Customer{
			Username:      "User Test",
			Password:      password,
			Amount:        50000,
			AccountNumber: "1234567890",
			AccountName:   "Account Name",
		}

		id, err := customerRepo.Create(ctx, arg2)

		require.NoError(t, err)
		require.NotEmpty(t, id)
	})

	t.Run("Get customer List", func(t *testing.T) {
		ctx := context.Background()

		customers, err := customerRepo.List(ctx)

		require.NoError(t, err)
		require.NotEmpty(t, customers)
	})

	t.Run("Get customer by username", func(t *testing.T) {
		ctx := context.Background()

		customer, err := customerRepo.GetByUsername(ctx, "User Test")

		require.NoError(t, err)
		require.NotEmpty(t, customer)
		require.Equal(t, customer.Username, "User Test")
	})

	t.Run("Get customer by account number", func(t *testing.T) {
		ctx := context.Background()

		customer, err := customerRepo.GetByAccountNumber(ctx, "1234567890")

		require.NoError(t, err)
		require.NotEmpty(t, customer)
		require.Equal(t, customer.AccountNumber, "1234567890")
	})

	t.Run("Get customer by id", func(t *testing.T) {
		ctx := context.Background()

		customer, err := customerRepo.GetByID(ctx, id)

		require.NoError(t, err)
		require.NotEmpty(t, customer)
		require.Equal(t, customer.ID, id)
	})

	t.Run("Update customer amount by id", func(t *testing.T) {
		ctx := context.Background()
		tx, err := TestDb.Begin()
		require.NoError(t, err)

		customer, err := customerRepo.GetByID(ctx, id)
		require.NoError(t, err)

		customer.Amount = 50000

		err = customerRepo.UpdateAmountByID(ctx, tx, customer)
		require.NoError(t, err)

		customer, err = customerRepo.GetByID(ctx, id)
		require.NoError(t, err)

		require.Equal(t, customer.Amount, 50000)
	})

	t.Run("Clean up data", func(t *testing.T) {
		clearData(t, TestDb)
	})
}

func TestCustomerRepository_Failed(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	customerRepo := NewCustomerRepository(TestDb)

	t.Run("Create customer with existing username", func(t *testing.T) {
		ctx := context.Background()
		password, err := utils.HashPassword("12345678")
		require.NoError(t, err)

		arg := model.Customer{
			Username:      "User Test",
			Password:      password,
			Amount:        50000,
			AccountNumber: "1234567890",
			AccountName:   "Account Name",
		}

		arg2 := model.Customer{
			Username:      "User Test",
			Password:      password,
			Amount:        50000,
			AccountNumber: "1234567890",
			AccountName:   "Account Name",
		}

		_, _ = customerRepo.Create(ctx, arg)
		_, err = customerRepo.Create(ctx, arg2)

		require.Error(t, err, "error is expected while creating customer with existing username")
		require.Contains(t, err.Error(), "duplicate key value violates unique constraint")
	})

	t.Run("Get non-existent customer by username", func(t *testing.T) {
		ctx := context.Background()

		_, err := customerRepo.GetByUsername(ctx, "Nonexistent User")

		require.Error(t, err, "error is expected while getting non-existent customer by username")
		require.EqualError(t, err, "sql: no rows in result set")
	})

	t.Run("List customers with empty table", func(t *testing.T) {
		ctx := context.Background()

		_, err := TestDb.Exec("DELETE FROM customers")
		require.NoError(t, err)

		customers, err := customerRepo.List(ctx)

		require.NoError(t, err, "error is not expected while listing customers from empty table")
		require.Empty(t, customers, "list of customers should be empty from an empty table")
	})

	t.Run("Get non-existent customer by account number", func(t *testing.T) {
		ctx := context.Background()

		_, err := customerRepo.GetByAccountNumber(ctx, "1234567890")

		require.Error(t, err, "error is expected while getting non-existent customer by account number")
		require.EqualError(t, err, "sql: no rows in result set")
	})

	t.Run("Get non-existent customer by id", func(t *testing.T) {
		ctx := context.Background()

		_, err := customerRepo.GetByID(ctx, uuid.New().String())

		require.Error(t, err, "error is expected while getting non-existent customer by id")
		require.EqualError(t, err, "sql: no rows in result set")
	})

	t.Run("Update non-existent customer amount by id", func(t *testing.T) {
		ctx := context.Background()
		tx, err := TestDb.Begin()
		require.NoError(t, err)

		customer := model.Customer{
			ID:     uuid.New().String(),
			Amount: 50000,
		}

		err = customerRepo.UpdateAmountByID(ctx, tx, customer)
		require.Error(t, err, "error is expected while updating non-existent customer amount by id")
		require.EqualError(t, err, "no rows updated")
	})

	t.Run("Clean up data", func(t *testing.T) {
		clearData(t, TestDb)
	})
}
