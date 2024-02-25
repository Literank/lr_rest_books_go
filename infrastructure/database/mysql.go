/*
Package database does all db persistence implementations.
*/
package database

import (
	"context"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"literank.com/rest-books/domain/model"
)

type MySQLPersistence struct {
	db *gorm.DB
}

func NewMySQLPersistence(dsn string) (*MySQLPersistence, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// Auto Migrate the data structs
	db.AutoMigrate(&model.Book{})

	return &MySQLPersistence{db}, nil
}

func (s *MySQLPersistence) CreateBook(ctx context.Context, b *model.Book) (uint, error) {
	if err := s.db.WithContext(ctx).Create(b).Error; err != nil {
		return 0, err
	}
	return b.ID, nil
}

func (s *MySQLPersistence) UpdateBook(ctx context.Context, id uint, b *model.Book) error {
	var book model.Book
	if err := s.db.WithContext(ctx).First(&book, id).Error; err != nil {
		return err
	}
	return s.db.WithContext(ctx).Model(book).Updates(b).Error
}

func (s *MySQLPersistence) DeleteBook(ctx context.Context, id uint) error {
	return s.db.WithContext(ctx).Delete(&model.Book{}, id).Error
}

func (s *MySQLPersistence) GetBook(ctx context.Context, id uint) (*model.Book, error) {
	var book model.Book
	if err := s.db.WithContext(ctx).First(&book, id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (s *MySQLPersistence) GetBooks(ctx context.Context) ([]*model.Book, error) {
	books := make([]*model.Book, 0)
	if err := s.db.WithContext(ctx).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}