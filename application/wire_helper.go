/*
Package application provides all common structures and functions of the application layer.
*/
package application

import (
	"literank.com/rest-books/domain/gateway"
	"literank.com/rest-books/infrastructure/cache"
	"literank.com/rest-books/infrastructure/config"
	"literank.com/rest-books/infrastructure/database"
	"literank.com/rest-books/infrastructure/token"
)

// WireHelper is the helper for dependency injection
type WireHelper struct {
	sqlPersistence   *database.MySQLPersistence
	noSQLPersistence *database.MongoPersistence
	kvStore          *cache.RedisCache
	tokenKeeper      *token.Keeper
}

// NewWireHelper constructs a new WireHelper
func NewWireHelper(c *config.Config) (*WireHelper, error) {
	db, err := database.NewMySQLPersistence(c.DB.DSN, c.App.PageSize)
	if err != nil {
		return nil, err
	}
	mdb, err := database.NewMongoPersistence(c.DB.MongoURI, c.DB.MongoDBName)
	if err != nil {
		return nil, err
	}
	kv := cache.NewRedisCache(&c.Cache)
	tk := token.NewTokenKeeper(c.App.TokenSecret, uint(c.App.TokenHours))
	return &WireHelper{
		sqlPersistence: db, noSQLPersistence: mdb,
		kvStore: kv, tokenKeeper: tk}, nil
}

// BookManager returns an instance of BookManager
func (w *WireHelper) BookManager() gateway.BookManager {
	return w.sqlPersistence
}

// UserManager returns an instance of UserManager
func (w *WireHelper) UserManager() gateway.UserManager {
	return w.sqlPersistence
}

// PermManager returns an instance of PermManager
func (w *WireHelper) PermManager() gateway.PermissionManager {
	return w.tokenKeeper
}

// ReviewManager returns an instance of ReviewManager
func (w *WireHelper) ReviewManager() gateway.ReviewManager {
	return w.noSQLPersistence
}

// CacheHelper returns an instance of CacheHelper
func (w *WireHelper) CacheHelper() cache.Helper {
	return w.kvStore
}
