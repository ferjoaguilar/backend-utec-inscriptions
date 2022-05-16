package database

import (
	"context"
	"errors"
	"time"

	"github.com/snowball-devs/backend-utec-inscriptions/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *MongodbRepository) CreateUser(ctx context.Context, user *models.User) (*mongo.InsertOneResult, error) {

	var findUser []models.User

	find, err := repo.DB.Collection("users").Find(ctx, bson.M{"permissions": "manager"})

	err = find.All(ctx, &findUser)
	if err != nil {
		return nil, err
	}

	if len(findUser) >= 3 {
		return nil, errors.New("You have exceeded the maximum number of users")
	}

	newUser := models.User{
		Email:       user.Email,
		Username:    user.Username,
		Password:    user.Password,
		Permissions: user.Permissions,
		Disable:     false,
		CreatedAt:   time.Now(),
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

func (repo *MongodbRepository) DisableUser(ctx context.Context, id string) (string, error) {

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}

	result := repo.DB.Collection("users").FindOneAndUpdate(ctx, bson.M{"_id": objId}, bson.D{
		{"$set", bson.D{{"disable", true}}},
	})

	if result.Err() != nil {
		return "", result.Err()
	}
	return "User updated successfully", nil

}

func (repo *MongodbRepository) GetManagers(ctx context.Context) ([]models.User, error) {
	result, err := repo.DB.Collection("users").Find(ctx, bson.M{"disable": false})
	if err != nil {
		return nil, err
	}
	var managerUsers []models.User
	err = result.All(ctx, &managerUsers)
	if err != nil {
		return nil, err
	}

	return managerUsers, nil

}
