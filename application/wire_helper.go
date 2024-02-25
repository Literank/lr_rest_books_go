package application

import (
	"literank.com/rest-books/domain/gateway"
	"literank.com/rest-books/infrastructure/config"
	"literank.com/rest-books/infrastructure/database"
)

// WireHelper is the helper for dependency injection
type WireHelper struct {
	sqlPersistence   *database.MySQLPersistence
	noSQLPersistence *database.MongoPersistence
}

func NewWireHelper(c *config.Config) (*WireHelper, error) {
	db, err := database.NewMySQLPersistence(c.DB.DSN)
	if err != nil {
		return nil, err
	}
	mdb, err := database.NewMongoPersistence(c.DB.MongoURI, c.DB.MongoDBName)
	if err != nil {
		return nil, err
	}
	return &WireHelper{sqlPersistence: db, noSQLPersistence: mdb}, nil
}

func (w *WireHelper) BookManager() gateway.BookManager {
	return w.sqlPersistence
}

func (w *WireHelper) ReviewManager() gateway.ReviewManager {
	return w.noSQLPersistence
}
