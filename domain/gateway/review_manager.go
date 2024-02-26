package gateway

import (
	"context"

	"literank.com/rest-books/domain/model"
)

// ReviewManager manages all book reviews
type ReviewManager interface {
	CreateReview(ctx context.Context, b *model.Review) (string, error)
	UpdateReview(ctx context.Context, id string, b *model.Review) error
	DeleteReview(ctx context.Context, id string) error
	GetReview(ctx context.Context, id string) (*model.Review, error)
	GetReviewsOfBook(ctx context.Context, bookID uint, keyword string) ([]*model.Review, error)
}
