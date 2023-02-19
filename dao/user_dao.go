package dao

import (
	"context"
	"github.com/NisalSP9/Did-I-read/commons"
	"github.com/NisalSP9/Did-I-read/connections"
	"github.com/NisalSP9/Did-I-read/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func CreateUser(user models.User) (models.User, *commons.RequestError) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	DB := connections.Connect()
	rst, err := DB.Collection("user").InsertOne(ctx, user)
	if err != nil {
		commons.ErrorLogger.Println(err.Error())
		connections.Disconnect(DB)
		return models.User{}, &commons.RequestError{StatusCode: http.StatusBadRequest, ErrorOccurredIn: "user_dao CreateUser", Err: err}
	}
	user.Id = rst.InsertedID.(primitive.ObjectID)
	connections.Disconnect(DB)
	return user, nil
}

func GetUserById(userId primitive.ObjectID) (models.User, *commons.RequestError) {
	var user models.User
	DB := connections.Connect()
	err := DB.Collection("user").FindOne(context.TODO(), bson.D{{"_id", userId}}).Decode(&user)
	if err != nil {
		commons.ErrorLogger.Println(err.Error())
		connections.Disconnect(DB)
		return models.User{}, &commons.RequestError{StatusCode: http.StatusBadRequest, ErrorOccurredIn: "user_dao GetUserById", Err: err}
	}
	connections.Disconnect(DB)
	return user, nil
}

func UserAuth(username, password string) (bool, *commons.RequestError) {
	var user models.User
	DB := connections.Connect()
	err := DB.Collection("user").FindOne(context.TODO(), bson.D{{"email", username}}).Decode(&user)
	if err != nil {
		commons.ErrorLogger.Println(err.Error())
		connections.Disconnect(DB)
		return false, &commons.RequestError{StatusCode: http.StatusBadRequest, ErrorOccurredIn: "user_dao UserAuth", Err: err}
	}
	connections.Disconnect(DB)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		commons.WarningLogger.Println(err.Error())
		return false, nil
	} else {
		return true, nil
	}
}
