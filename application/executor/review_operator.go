package executor

import (
	"context"
	"errors"
	"time"

	"literank.com/rest-books/application/dto"
	"literank.com/rest-books/domain/gateway"
	"literank.com/rest-books/domain/model"
)

type ReviewOperator struct {
	reviewManager gateway.ReviewManager
}

func NewReviewOperator(b gateway.ReviewManager) *ReviewOperator {
	return &ReviewOperator{reviewManager: b}
}

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

func (o *ReviewOperator) GetReview(ctx context.Context, id string) (*model.Review, error) {
	return o.reviewManager.GetReview(ctx, id)
}

func (o *ReviewOperator) GetReviewsOfBook(ctx context.Context, bookID uint) ([]*model.Review, error) {
	return o.reviewManager.GetReviewsOfBook(ctx, bookID)
}

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

func (o *ReviewOperator) DeleteReview(ctx context.Context, id string) error {
	return o.reviewManager.DeleteReview(ctx, id)
}
