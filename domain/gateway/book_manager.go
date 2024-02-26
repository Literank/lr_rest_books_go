package gateway

import (
	"context"

	"literank.com/rest-books/domain/model"
)

// BookManager manages all books
type BookManager interface {
	CreateBook(ctx context.Context, b *model.Book) (uint, error)
	UpdateBook(ctx context.Context, id uint, b *model.Book) error
	DeleteBook(ctx context.Context, id uint) error
	GetBook(ctx context.Context, id uint) (*model.Book, error)
	GetBooks(ctx context.Context, offset int, keyword string) ([]*model.Book, error)
}
