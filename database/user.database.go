package database

import (
	"context"
	"time"

	"github.com/snowball-devs/backend-utec-inscriptions/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *MongodbRepository) CreateUser(ctx context.Context, user *models.User) (*mongo.InsertOneResult, error) {
	newUser := models.User{
		Email:     user.Email,
		Username:  user.Username,
		Password:  user.Password,
		Disable:   false,
		CreatedAt: time.Now(),
	}

	result, err := repo.DB.Collection("users").InsertOne(ctx, newUser)

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (repo *MongodbRepository) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user *models.User

	err := repo.DB.Collection("users").FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
