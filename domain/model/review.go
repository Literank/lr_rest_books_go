package model

import "time"

// Review represents the review of a book
type Review struct {
	// Caution: bson tag, which is bound to a specific db, should not appear here in the domain entity.
	// It's a hack for tutorial brevity.
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	BookID    uint      `json:"book_id,omitempty"`
	Author    string    `json:"author,omitempty"`
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
