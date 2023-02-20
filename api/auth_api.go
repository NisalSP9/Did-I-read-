package api

import (
	"encoding/json"
	"github.com/NisalSP9/Did-I-read/commons"
	"github.com/NisalSP9/Did-I-read/controllers"
	"net/http"
)

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
		w.WriteHeader(err1.StatusCode)
		if err := json.NewEncoder(w).Encode(err1.Err); err != nil {
			commons.ErrorLogger.Println(err1.Err)
		}
	} else {
		// TODO: add a log
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(auth); err != nil {
			commons.ErrorLogger.Println(err)
		}
	}
}

func RefreshToke(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	refreshedToken, err := commons.RefreshToken(r.Header.Get("Authorization"))

	if err != nil {
		commons.ErrorLogger.Println(err.Err)
		w.WriteHeader(err.StatusCode)
		if err := json.NewEncoder(w).Encode(err.Err); err != nil {
			commons.ErrorLogger.Println(err)
			return
		}
	} else {
		// TODO: add a log
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(refreshedToken); err != nil {
			commons.ErrorLogger.Println(err)
		}
	}

}
