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

// MySQLPersistence runs all MySQL operations
type MySQLPersistence struct {
	db       *gorm.DB
	pageSize int
}

// NewMySQLPersistence constructs a new MySQLPersistence
func NewMySQLPersistence(dsn string, pageSize int) (*MySQLPersistence, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// Auto Migrate the data structs
	if err := db.AutoMigrate(&model.Book{}, &model.User{}); err != nil {
		return nil, err
	}

	return &MySQLPersistence{db, pageSize}, nil
}

// CreateBook creates a new book
func (s *MySQLPersistence) CreateBook(ctx context.Context, b *model.Book) (uint, error) {
	if err := s.db.WithContext(ctx).Create(b).Error; err != nil {
		return 0, err
	}
	return b.ID, nil
}

// UpdateBook updates a book by its ID and the new content
func (s *MySQLPersistence) UpdateBook(ctx context.Context, id uint, b *model.Book) error {
	var book model.Book
	if err := s.db.WithContext(ctx).First(&book, id).Error; err != nil {
		return err
	}
	return s.db.WithContext(ctx).Model(book).Updates(b).Error
}

// DeleteBook deletes a book by ID
func (s *MySQLPersistence) DeleteBook(ctx context.Context, id uint) error {
	return s.db.WithContext(ctx).Delete(&model.Book{}, id).Error
}

// GetBook gets a book by ID
func (s *MySQLPersistence) GetBook(ctx context.Context, id uint) (*model.Book, error) {
	var book model.Book
	if err := s.db.WithContext(ctx).First(&book, id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

// GetBooks gets a list of books by offset and keyword
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

// CreateUser creates a new user
func (s *MySQLPersistence) CreateUser(ctx context.Context, u *model.User) (uint, error) {
	if err := s.db.WithContext(ctx).Create(u).Error; err != nil {
		return 0, err
	}
	return u.ID, nil
}

// GetUserByEmail gets the user by its email
func (s *MySQLPersistence) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var u model.User
	if err := s.db.WithContext(ctx).Where("email = ?", email).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}
