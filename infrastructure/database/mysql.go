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
	db       *gorm.DB
	pageSize int
}

func NewMySQLPersistence(dsn string, pageSize int) (*MySQLPersistence, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// Auto Migrate the data structs
	db.AutoMigrate(&model.Book{})

	return &MySQLPersistence{db, pageSize}, nil
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

func (s *MySQLPersistence) GetBooks(ctx context.Context, offset int, keyword string) ([]*model.Book, error) {
	books := make([]*model.Book, 0)
	tx := s.db.WithContext(ctx)
	if keyword != "" {
		term := "%" + keyword + "%"
		tx = tx.Where("title LIKE ?", term).Or("author LIKE ?", term)
	}
	if err := tx.Offset(offset).Limit(s.pageSize).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}
