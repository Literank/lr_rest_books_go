/*
Package database does all db persistence implementations.
*/
package database

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"literank.com/rest-books/domain/model"
)

const (
	collReview  = "reviews"
	idField     = "_id"
	bookIDField = "bookid"
)

type MongoPersistence struct {
	db   *mongo.Database
	coll *mongo.Collection
}

func NewMongoPersistence(mongoURI, dbName string) (*MongoPersistence, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}
	db := client.Database(dbName)
	coll := db.Collection(collReview)
	return &MongoPersistence{db, coll}, nil
}

func (m *MongoPersistence) CreateReview(ctx context.Context, r *model.Review) (string, error) {
	result, err := m.coll.InsertOne(ctx, r)
	if err != nil {
		return "", err
	}
	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("failed to extract InsertedID from result: %v", result)
	}
	return insertedID.Hex(), nil
}

func (m *MongoPersistence) UpdateReview(ctx context.Context, id string, r *model.Review) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	updateValues := bson.M{
		"title":     r.Title,
		"content":   r.Content,
		"updatedat": r.UpdatedAt,
	}
	result, err := m.coll.UpdateOne(ctx, bson.M{idField: objID}, bson.M{"$set": updateValues})
	if err != nil {
		return err
	}
	if result.ModifiedCount == 0 {
		return errors.New("review does not exist")
	}
	return nil
}

func (m *MongoPersistence) DeleteReview(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	result, err := m.coll.DeleteOne(ctx, bson.M{idField: objID})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("review does not exist")
	}
	return nil
}

func (m *MongoPersistence) GetReview(ctx context.Context, id string) (*model.Review, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var review model.Review
	if err := m.coll.FindOne(ctx, bson.M{idField: objID}).Decode(&review); err != nil {
		return nil, err
	}
	return &review, nil
}

func (m *MongoPersistence) GetReviewsOfBook(ctx context.Context, bookID uint) ([]*model.Review, error) {
	cursor, err := m.coll.Find(ctx, bson.M{bookIDField: bookID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	reviews := make([]*model.Review, 0)
	if err := cursor.All(ctx, &reviews); err != nil {
		return nil, err
	}
	return reviews, nil
}
