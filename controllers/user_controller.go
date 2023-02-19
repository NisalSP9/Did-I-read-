package controllers

import (
	"github.com/NisalSP9/Did-I-read/commons"
	"github.com/NisalSP9/Did-I-read/dao"
	"github.com/NisalSP9/Did-I-read/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func CreateUser(user models.User) (models.User, *commons.RequestError) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, &commons.RequestError{StatusCode: http.StatusBadRequest, ErrorOccurredIn: "user_controller CreateUser", Err: err}
	}
	user.Password = string(hashedPassword)
	return dao.CreateUser(user)
}

func GetUserById(userId primitive.ObjectID) (models.User, *commons.RequestError) {
	return dao.GetUserById(userId)
}

func UserAuth(username, password string) (bool, *commons.RequestError) {
	return dao.UserAuth(username, password)
}
