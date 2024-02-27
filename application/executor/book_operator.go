/*
Package executor handles request-response style business logic.
*/
package executor

import (
	"context"
	"encoding/json"
	"fmt"

	"literank.com/rest-books/domain/gateway"
	"literank.com/rest-books/domain/model"
	"literank.com/rest-books/infrastructure/cache"
)

const booksKey = "lr-books"

// BookOperator handles book input/output and proxies operations to the book manager.
type BookOperator struct {
	bookManager gateway.BookManager
	cacheHelper cache.Helper
}

// NewBookOperator constructs a new BookOperator
func NewBookOperator(b gateway.BookManager, c cache.Helper) *BookOperator {
	return &BookOperator{bookManager: b, cacheHelper: c}
}

// CreateBook creates a new book
func (o *BookOperator) CreateBook(ctx context.Context, b *model.Book) (*model.Book, error) {
	id, err := o.bookManager.CreateBook(ctx, b)
	if err != nil {
		return nil, err
	}
	b.ID = id
	return b, nil
}

// GetBook gets a book by ID
func (o *BookOperator) GetBook(ctx context.Context, id uint) (*model.Book, error) {
	return o.bookManager.GetBook(ctx, id)
}

// GetBooks gets a list of books by offset and keyword, and caches its result if needed
func (o *BookOperator) GetBooks(ctx context.Context, offset int, query string) ([]*model.Book, error) {
	// Search results, don't cache it
	if query != "" {
		return o.bookManager.GetBooks(ctx, offset, query)
	}

	// Normal list of results
	k := fmt.Sprintf("%s-%d", booksKey, offset)
	rawValue, err := o.cacheHelper.Load(ctx, k)
	if err != nil {
		return nil, err
	}

	books := make([]*model.Book, 0)
	if rawValue != "" {
		// Cache key exists
		if err = json.Unmarshal([]byte(rawValue), &books); err != nil {
			return nil, err
		}
	} else {
		// Cache key does not exist
		books, err = o.bookManager.GetBooks(ctx, offset, "")
		if err != nil {
			return nil, err
		}
		value, err := json.Marshal(books)
		if err != nil {
			return nil, err
		}
		if err := o.cacheHelper.Save(ctx, k, string(value)); err != nil {
			return nil, err
		}
	}
	return books, nil
}

// UpdateBook updates a book by its ID and the new content
func (o *BookOperator) UpdateBook(ctx context.Context, id uint, b *model.Book) (*model.Book, error) {
	if err := o.bookManager.UpdateBook(ctx, id, b); err != nil {
		return nil, err
	}
	return b, nil
}

// DeleteBook deletes a book by ID
func (o *BookOperator) DeleteBook(ctx context.Context, id uint) error {
	return o.bookManager.DeleteBook(ctx, id)
}
