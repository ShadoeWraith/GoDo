package common

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

func JSONResponse(w http.ResponseWriter, data interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		JSONError(w, http.StatusInternalServerError, "Error encoding JSON")
		return
	}
}

func JSONSuccess(w http.ResponseWriter, data interface{}, status int) {
	JSONResponse(w, data, status)
}

func JSONError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	response := ErrorResponse{
		Status: status,
		Error:  message,
	}

	json.NewEncoder(w).Encode(response)
}
