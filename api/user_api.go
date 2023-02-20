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

	err1 := controllers.CreateUser(&user)
	if err1 != nil {
		commons.ErrorLogger.Println(err1)
		w.WriteHeader(err1.StatusCode)
		if err := json.NewEncoder(w).Encode(err1.Err); err != nil {
			commons.ErrorLogger.Println(err)
		}
	} else {
		// TODO: add a log
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(user); err != nil {
			commons.ErrorLogger.Println(err)
		}
	}
}

func GetUserById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	status, _, authErr := commons.IsAuthorized(r.Header.Get("Authorization"))

	if authErr != nil {
		commons.ErrorLogger.Println(authErr.Err)
		w.WriteHeader(authErr.StatusCode)
		if err := json.NewEncoder(w).Encode(authErr.Err); err != nil {
			commons.ErrorLogger.Println(err)
		}
		return
	}

	if status {
		vars := mux.Vars(r)
		userID := vars["userid"]
		objectId, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			commons.ErrorLogger.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			if err := json.NewEncoder(w).Encode(err.Error()); err != nil {
				commons.ErrorLogger.Println(err)
			}
			return
		}

		foundUser, err1 := controllers.GetUserById(objectId)
		if err1 != nil {
			commons.ErrorLogger.Println(err1)
			w.WriteHeader(err1.StatusCode)
			if err := json.NewEncoder(w).Encode(err1.Err); err != nil {
				commons.ErrorLogger.Println(err)
			}
			return
		} else {
			// TODO: add a log
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(foundUser); err != nil {
				commons.ErrorLogger.Println(err)
				return
			}
		}
	} else {
		commons.ErrorLogger.Println("Unauthorized")
		w.WriteHeader(http.StatusUnauthorized)
		if err := json.NewEncoder(w).Encode("Unauthorized"); err != nil {
			commons.ErrorLogger.Println(err)
			return
		}
	}

}
