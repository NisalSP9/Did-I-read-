package dao

import (
	"context"
	"net/http"
	"time"

	"github.com/NisalSP9/Did-I-read/commons"
	"github.com/NisalSP9/Did-I-read/connections"
	"github.com/NisalSP9/Did-I-read/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(user *models.User) *commons.RequestError {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rst, err := connections.DB.Collection("user").InsertOne(ctx, *user)
	if err != nil {
		commons.ErrorLogger.Println(err.Error())
		return &commons.RequestError{StatusCode: http.StatusBadRequest, ErrorOccurredIn: "user_dao CreateUser", Err: err.Error()}
	}
	user.Id = rst.InsertedID.(primitive.ObjectID)
	return nil
}

func GetUserById(userId primitive.ObjectID) (*models.User, *commons.RequestError) {
	var user models.User
	err := connections.DB.Collection("user").FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: userId}}).Decode(&user)
	if err != nil {
		commons.ErrorLogger.Println(err.Error())
		return &models.User{}, &commons.RequestError{StatusCode: http.StatusBadRequest, ErrorOccurredIn: "user_dao GetUserById", Err: err.Error()}
	}
	return &user, nil
}

func UserAuth(username, password string) (string, *commons.RequestError) {
	var user models.User
	err := connections.DB.Collection("user").FindOne(context.TODO(), bson.D{primitive.E{Key: "email", Value: username}}).Decode(&user)
	if err != nil {
		commons.ErrorLogger.Println(err.Error())
		return "", &commons.RequestError{StatusCode: http.StatusBadRequest, ErrorOccurredIn: "user_dao UserAuth", Err: err.Error()}
	}
	return commons.GetAuthToken(username, password, user.Password)
}
