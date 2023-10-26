package bank

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/lib/pq"
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

func insertData(t *testing.T, db *sql.DB) {
	_, err := db.Exec("INSERT INTO banks (name, account_number, account_name) VALUES ('BNI', '1234567890', 'Admin'), ('BCA', '0987654321', 'Admin 2')")
	if err != nil {
		t.Fatalf("failed to clear data: %v", err)
	}
}

func TestBanRepository_Success(t *testing.T) {
	if testing.Short() {
		os.Exit(0)
	}

	bankRepo := NewBankRepository(TestDb)

	t.Run("List bank", func(t *testing.T) {
		ctx := context.Background()

		banks, err := bankRepo.List(ctx)

		require.NoError(t, err)
		require.NotEmpty(t, banks)
	})

	t.Run("Get bank by account number", func(t *testing.T) {
		ctx := context.Background()

		bank, err := bankRepo.GetByAccountNumber(ctx, "1234567890")

		require.NoError(t, err)
		require.NotEmpty(t, bank)
	})
}

func TestBankRepository_Failed(t *testing.T) {
	if testing.Short() {
		os.Exit(0)
	}

	bankRepo := NewBankRepository(TestDb)

	t.Run("List banks with empty table", func(t *testing.T) {
		ctx := context.Background()

		_, err := TestDb.Exec("DELETE FROM banks")
		require.NoError(t, err)

		banks, err := bankRepo.List(ctx)

		require.NoError(t, err, "error is not expected while listing banks from empty table")
		require.Empty(t, banks, "list of banks should be empty from an empty table")
	})

	t.Run("Get bank by account number with empty table", func(t *testing.T) {
		ctx := context.Background()

		_, err := TestDb.Exec("DELETE FROM banks")
		require.NoError(t, err)

		bank, err := bankRepo.GetByAccountNumber(ctx, "1234567890")

		require.Error(t, err, "error is expected while getting bank from empty table")
		require.Empty(t, bank, "bank should be empty from an empty table")
	})

	t.Run("Insert up data", func(t *testing.T) {
		insertData(t, TestDb)
	})
}
