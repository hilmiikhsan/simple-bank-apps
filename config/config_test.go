package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	t.Run("Load env success", func(t *testing.T) {
		err := LoadConfig("../env.yaml")

		expectedSource := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			Cfg.DB.Host, Cfg.DB.Port, Cfg.DB.User, Cfg.DB.Password, Cfg.DB.Name,
		)
		expectedSourceTest := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			"localhost", "5432", "postgres", "21012123op", "bank_apps",
		)
		expectedDriver := "postgres"
		expectedAddress := ":4000"

		require.NoError(t, err)
		require.NotEmpty(t, Cfg)
		require.Equal(t, expectedSourceTest, expectedSource)
		require.Equal(t, expectedDriver, Cfg.DB.Driver)
		require.Equal(t, expectedAddress, Cfg.App.Port)
	})

	t.Run("Load env fail", func(t *testing.T) {
		err := LoadConfig("./env.yaml")

		require.Error(t, err)
	})
}
