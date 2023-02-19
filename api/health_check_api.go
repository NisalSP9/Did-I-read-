package api

import (
	"encoding/json"
	"github.com/NisalSP9/Did-I-read/commons"
	"net/http"
)

type HealthResponse struct {
	Status bool   `json:"status"`
	Name   string `json:"name"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {

	resp := HealthResponse{
		Status: true,
		Name:   "Did i read admin up and running",
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		commons.ErrorLogger.Println("Error: %v\n", err.Error())
	}
	return

}
