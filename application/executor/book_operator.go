package executor

import (
	"context"

	"literank.com/rest-books/domain/gateway"
	"literank.com/rest-books/domain/model"
)

type BookOperator struct {
	bookManager gateway.BookManager
}

func NewBookOperator(b gateway.BookManager) *BookOperator {
	return &BookOperator{bookManager: b}
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
	return o.bookManager.GetBooks(ctx)
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
