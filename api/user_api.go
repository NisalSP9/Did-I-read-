package api

import (
	"encoding/json"
	"github.com/NisalSP9/Did-I-read/commons"
	"github.com/NisalSP9/Did-I-read/controllers"
	"github.com/NisalSP9/Did-I-read/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	decoder := json.NewDecoder(r.Body)
	var user models.User
	err := decoder.Decode(&user)
	if err != nil {
		commons.ErrorLogger.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			commons.ErrorLogger.Println(err)
		}
	}

	createdUser, err1 := controllers.CreateUser(user)
	if err1 != nil {
		commons.ErrorLogger.Println(err1)
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(err1); err != nil {
			commons.ErrorLogger.Println(err)
		}
	} else {
		// TODO: add a log
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(createdUser); err != nil {
			commons.ErrorLogger.Println(err)
		}
	}
}

func GetUserById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	userID := vars["userid"]
	objectId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		commons.ErrorLogger.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			commons.ErrorLogger.Println(err)
		}
	}

	foundUser, err1 := controllers.GetUserById(objectId)
	if err1 != nil {
		commons.ErrorLogger.Println(err1)
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(err1); err != nil {
			commons.ErrorLogger.Println(err)
		}
	} else {
		// TODO: add a log
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(foundUser); err != nil {
			commons.ErrorLogger.Println(err)
		}
	}
}

type LoginRequestWrapper struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func UserAuth(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	decoder := json.NewDecoder(r.Body)
	var loginRequestWrapper LoginRequestWrapper
	err := decoder.Decode(&loginRequestWrapper)
	if err != nil {
		commons.ErrorLogger.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			commons.ErrorLogger.Println(err)
		}
	}

	auth, err1 := controllers.UserAuth(loginRequestWrapper.Username, loginRequestWrapper.Password)
	if err1 != nil {
		commons.ErrorLogger.Println(err1)
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(false); err != nil {
			commons.ErrorLogger.Println(err)
		}
	} else {
		// TODO: add a log
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(auth); err != nil {
			commons.ErrorLogger.Println(err)
		}
	}
}
