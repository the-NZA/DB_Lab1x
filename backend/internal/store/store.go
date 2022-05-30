package store

import (
	"errors"

	"github.com/the-NZA/DB_Lab1x/backend/internal/config"
	"github.com/the-NZA/DB_Lab1x/backend/internal/store/mock"
	"github.com/the-NZA/DB_Lab1x/backend/internal/store/mysql"
	"github.com/the-NZA/DB_Lab1x/backend/internal/store/postgres"
	"github.com/the-NZA/DB_Lab1x/backend/internal/store/sqlite3"
	"github.com/the-NZA/DB_Lab1x/backend/internal/store/storer"
)

var (
	ErrUnsupportedDBType = errors.New("unsupported DB type")
)

func NewStore(config *config.Config) (storer.Storer, error) {
	switch config.DBType {
	case "mysql":
		return mysql.NewStore(config)
	case "sqlite3":
		return sqlite3.NewStore(config)
	case "mock":
		return mock.NewStore(config)
	case "postgres":
		return postgres.NewStore(config)
	default:
		return nil, ErrUnsupportedDBType
	}
}
