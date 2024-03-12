package api

import (
	"encoding/json"
	"net/http"

	"github.com/NisalSP9/Did-I-read/commons"
	"github.com/NisalSP9/Did-I-read/connections"
)

type HealthResponse struct {
	Status bool   `json:"status"`
	Name   string `json:"name"`
	DatabaseConnection bool `json:"databaseConnection"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {

	rst := connections.CheckConnection()

	resp := HealthResponse{
		Status: true,
		Name:   "Did i read admin up and running",
		DatabaseConnection: rst,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		commons.ErrorLogger.Println("Error: %v\n", err.Error())
	}
	return

}
