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

func (w *WireHelper) BookManager() gateway.BookManager {
	return w.sqlPersistence
}

func (w *WireHelper) UserManager() gateway.UserManager {
	return w.sqlPersistence
}

func (w *WireHelper) PermManager() gateway.PermissionManager {
	return w.tokenKeeper
}

func (w *WireHelper) ReviewManager() gateway.ReviewManager {
	return w.noSQLPersistence
}

func (w *WireHelper) CacheHelper() cache.Helper {
	return w.kvStore
}
