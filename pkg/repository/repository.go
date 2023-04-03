package repository

import (
	"context"
	"mage/pkg/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{db.Collection("users")}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	user.CreatedAt = time.Now()
	counter, _ := r.CountAllUsers(ctx)
	user.UserID = counter + 1
	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *models.User) error {
	filter := bson.M{"_id": user.Id}
	update := bson.M{"$set": bson.M{"username": user.Username, "password": user.Password}}
	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *UserRepository) DeleteUser(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := r.collection.DeleteOne(ctx, filter)
	return err
}

func (r *UserRepository) GetUserById(ctx context.Context, id primitive.ObjectID) (*models.User, error) {
	filter := bson.M{"_id": id}
	var user models.User
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	filter := bson.M{"username": username}
	var user models.User
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CountAllUsers(ctx context.Context) (int64, error) {
	count, err := r.collection.EstimatedDocumentCount(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}
