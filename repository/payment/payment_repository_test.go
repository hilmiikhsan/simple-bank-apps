package payment

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/simple-bank-apps/model"
	"github.com/simple-bank-apps/repository/customer"
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

func createRandomCustomer(t *testing.T, customerRepo customer.CustomerRepository) {
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
	_, err := db.Exec("DELETE FROM payments WHERE customer_id = (SELECT id FROM customers WHERE username = 'User Test')")
	if err != nil {
		t.Fatalf("failed to clear data: %v", err)
	}

	_, err = db.Exec("DELETE FROM customers WHERE username = 'User Test'")
	if err != nil {
		t.Fatalf("failed to clear data: %v", err)
	}
}

func TestPaymentRepositorySuccess(t *testing.T) {
	ctx := context.Background()
	customerRepo := customer.NewCustomerRepository(TestDb)
	paymentRepo := NewPaymentRepository(TestDb)

	createRandomCustomer(t, customerRepo)

	customer, err := customerRepo.GetByUsername(ctx, "User Test")
	require.NoError(t, err)

	t.Run("Create payment success", func(t *testing.T) {
		arg := model.Payment{
			CustomerID: customer.ID,
			Amount:     100000,
		}

		err := paymentRepo.Create(ctx, arg)

		require.NoError(t, err)
		require.Equal(t, arg.CustomerID, customer.ID)
		require.Equal(t, arg.Amount, 100000)
	})

	t.Run("Clean up data", func(t *testing.T) {
		clearData(t, TestDb)
	})
}

func TestPaymentRepositoryFailed(t *testing.T) {
	ctx := context.Background()
	paymentRepo := NewPaymentRepository(TestDb)

	t.Run("Create payment fail", func(t *testing.T) {
		arg := model.Payment{
			CustomerID: "sadad",
			Amount:     100000,
		}

		err := paymentRepo.Create(ctx, arg)

		require.Error(t, err)
	})
}
