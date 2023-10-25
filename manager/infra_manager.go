package manager

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/simple-bank-apps/config"
)

type InfraManager interface {
	Connect() *sql.DB
	Config() *config.Config
}

type infraManager struct {
	db     *sql.DB
	config *config.Config
}

func (i *infraManager) initDb() error {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Cfg.DB.Host, config.Cfg.DB.Port, config.Cfg.DB.User, config.Cfg.DB.Password, config.Cfg.DB.Name,
	)

	db, err := sql.Open(config.Cfg.DB.Driver, dsn)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	i.db = db

	return nil
}

func (i *infraManager) Connect() *sql.DB {
	return i.db
}

func (i *infraManager) Config() *config.Config {
	return i.config
}

func NewInfraManager(configParam *config.Config) (InfraManager, error) {
	infra := &infraManager{
		config: configParam,
	}

	err := infra.initDb()
	if err != nil {
		return nil, err
	}

	return infra, nil
}
