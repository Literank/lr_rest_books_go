package executor

import (
	"context"
	"encoding/json"

	"literank.com/rest-books/domain/gateway"
	"literank.com/rest-books/domain/model"
	"literank.com/rest-books/infrastructure/cache"
)

const booksKey = "lr-books"

type BookOperator struct {
	bookManager gateway.BookManager
	cacheHelper cache.Helper
}

func NewBookOperator(b gateway.BookManager, c cache.Helper) *BookOperator {
	return &BookOperator{bookManager: b, cacheHelper: c}
}

func (o *BookOperator) CreateBook(ctx context.Context, b *model.Book) (*model.Book, error) {
	id, err := o.bookManager.CreateBook(ctx, b)
	if err != nil {
		return nil, err
	}
	b.ID = id
	return b, nil
}

func (o *BookOperator) GetBook(ctx context.Context, id uint) (*model.Book, error) {
	return o.bookManager.GetBook(ctx, id)
}

func (o *BookOperator) GetBooks(ctx context.Context) ([]*model.Book, error) {
	rawValue, err := o.cacheHelper.Load(ctx, booksKey)
	if err != nil {
		return nil, err
	}

	books := make([]*model.Book, 0)
	if rawValue != "" {
		// Cache key exists
		if err := json.Unmarshal([]byte(rawValue), &books); err != nil {
			return nil, err
		}
	} else {
		// Cache key does not exist
		books, err = o.bookManager.GetBooks(ctx)
		if err != nil {
			return nil, err
		}
		value, err := json.Marshal(books)
		if err != nil {
			return nil, err
		}
		if err := o.cacheHelper.Save(ctx, booksKey, string(value)); err != nil {
			return nil, err
		}
	}
	return books, nil
}

func (o *BookOperator) UpdateBook(ctx context.Context, id uint, b *model.Book) (*model.Book, error) {
	if err := o.bookManager.UpdateBook(ctx, id, b); err != nil {
		return nil, err
	}
	return b, nil
}

func (o *BookOperator) DeleteBook(ctx context.Context, id uint) error {
	return o.bookManager.DeleteBook(ctx, id)
}
