package model

import "time"

// Book represents the structure of a book
type Book struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	PublishedAt time.Time `json:"published_at"`
	Description string    `json:"description"`
	ISBN        string    `json:"isbn"`
	TotalPages  int       `json:"total_pages"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
