/*
Package database does all db persistence implementations.
*/
package database

import (
	"context"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"literank.com/rest-books/domain/model"
)

// SQLitePersistence runs all SQLite operations
type SQLitePersistence struct {
	db *gorm.DB
}

// NewSQLitePersistence constructs a new SQLitePersistence
func NewSQLitePersistence(fileName string) (*SQLitePersistence, error) {
	db, err := gorm.Open(sqlite.Open(fileName), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &SQLitePersistence{db}, nil
}

// CreateBook creates a new book
func (s *SQLitePersistence) CreateBook(ctx context.Context, b *model.Book) (uint, error) {
	if err := s.db.WithContext(ctx).Create(b).Error; err != nil {
		return 0, err
	}
	return b.ID, nil
}

// UpdateBook updates a book by its ID and the new content
func (s *SQLitePersistence) UpdateBook(ctx context.Context, id uint, b *model.Book) error {
	var book model.Book
	if err := s.db.WithContext(ctx).First(&book, id).Error; err != nil {
		return err
	}
	return s.db.WithContext(ctx).Model(book).Updates(b).Error
}

// DeleteBook deletes a book by ID
func (s *SQLitePersistence) DeleteBook(ctx context.Context, id uint) error {
	return s.db.WithContext(ctx).Delete(&model.Book{}, id).Error
}

// GetBook gets a book by ID
func (s *SQLitePersistence) GetBook(ctx context.Context, id uint) (*model.Book, error) {
	var book model.Book
	if err := s.db.WithContext(ctx).First(&book, id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

// GetBooks gets all books
func (s *SQLitePersistence) GetBooks(ctx context.Context) ([]*model.Book, error) {
	books := make([]*model.Book, 0)
	if err := s.db.WithContext(ctx).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}
