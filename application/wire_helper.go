package application

import (
	"literank.com/rest-books/domain/gateway"
	"literank.com/rest-books/infrastructure/config"
	"literank.com/rest-books/infrastructure/database"
)

// WireHelper is the helper for dependency injection
type WireHelper struct {
	persistence *database.MySQLPersistence
}

func NewWireHelper(c *config.Config) (*WireHelper, error) {
	db, err := database.NewMySQLPersistence(c.DB.DSN)
	if err != nil {
		return nil, err
	}
	return &WireHelper{persistence: db}, nil
}

func (w *WireHelper) BookManager() gateway.BookManager {
	return w.persistence
}
