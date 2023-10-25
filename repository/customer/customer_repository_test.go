package customer

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

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

func createRandomCustomer(t *testing.T, customerRepo CustomerRepository) {
	ctx := context.Background()
	password, err := utils.HashPassword("12345678")
	require.NoError(t, err)

	arg := model.Customer{
		Username: "User Test",
		Password: password,
	}

	err = customerRepo.Create(ctx, arg)

	require.NoError(t, err)
	require.Equal(t, arg.Username, "User Test")
	require.Equal(t, arg.Password, password)
}

func clearData(t *testing.T, db *sql.DB) {
	_, err := db.Exec("DELETE FROM customers WHERE username = 'User Test'")
	if err != nil {
		t.Fatalf("failed to clear data: %v", err)
	}
}

func TestCustomerRepositorySuccess(t *testing.T) {
	if testing.Short() {
		os.Exit(0)
	}

	customerRepo := NewCustomerRepository(TestDb)

	t.Run("Create customer", func(t *testing.T) {
		createRandomCustomer(t, customerRepo)
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

	t.Run("Clean up data", func(t *testing.T) {
		clearData(t, TestDb)
	})
}

func TestCustomerRepositoryFailed(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	customerRepo := NewCustomerRepository(TestDb)

	t.Run("Create customer with existing username", func(t *testing.T) {
		ctx := context.Background()
		password, err := utils.HashPassword("12345678")
		require.NoError(t, err)

		arg := model.Customer{
			Username: "User Test",
			Password: password,
		}

		arg2 := model.Customer{
			Username: "User Test",
			Password: password,
		}

		_ = customerRepo.Create(ctx, arg)
		err = customerRepo.Create(ctx, arg2)

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

	t.Run("Clean up data", func(t *testing.T) {
		clearData(t, TestDb)
	})
}
