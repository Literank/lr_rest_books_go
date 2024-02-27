package executor

import (
	"context"
	"errors"
	"time"

	"literank.com/rest-books/application/dto"
	"literank.com/rest-books/domain/gateway"
	"literank.com/rest-books/domain/model"
)

// ReviewOperator handles review input/output and proxies operations to the review manager.
type ReviewOperator struct {
	reviewManager gateway.ReviewManager
}

// NewReviewOperator constructs a new ReviewOperator
func NewReviewOperator(b gateway.ReviewManager) *ReviewOperator {
	return &ReviewOperator{reviewManager: b}
}

// CreateReview creates a new review
func (o *ReviewOperator) CreateReview(ctx context.Context, body *dto.ReviewBody) (*model.Review, error) {
	now := time.Now()
	b := &model.Review{
		BookID:    body.BookID,
		Author:    body.Author,
		Title:     body.Title,
		Content:   body.Content,
		CreatedAt: now,
		UpdatedAt: now,
	}
	id, err := o.reviewManager.CreateReview(ctx, b)
	if err != nil {
		return nil, err
	}
	b.ID = id
	return b, nil
}

// GetReview gets a review by ID
func (o *ReviewOperator) GetReview(ctx context.Context, id string) (*model.Review, error) {
	return o.reviewManager.GetReview(ctx, id)
}

// GetReviewsOfBook gets a list of reviews by a query
func (o *ReviewOperator) GetReviewsOfBook(ctx context.Context, bookID uint, query string) ([]*model.Review, error) {
	return o.reviewManager.GetReviewsOfBook(ctx, bookID, query)
}

// UpdateReview updates a review by its ID and the new content
func (o *ReviewOperator) UpdateReview(ctx context.Context, id string, b *model.Review) (*model.Review, error) {
	if b.Title == "" || b.Content == "" {
		return nil, errors.New("required field cannot be empty")
	}
	b.UpdatedAt = time.Now()
	if err := o.reviewManager.UpdateReview(ctx, id, b); err != nil {
		return nil, err
	}
	return b, nil
}

// DeleteReview deletes a review by ID
func (o *ReviewOperator) DeleteReview(ctx context.Context, id string) error {
	return o.reviewManager.DeleteReview(ctx, id)
}
