package common

import (
	"encoding/json"
	"net/http"
)

type ResponseError struct {
	Error ErrorDetail `json:"error"`
}

type ErrorDetail struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewResponseError(code int, message string) ResponseError {
	return ResponseError{
		Error: ErrorDetail{
			Code:    code,
			Message: message,
		},
	}
}

func InternalServerError(w http.ResponseWriter, r *http.Request) {
	if r := recover(); r != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		responseError := NewResponseError(http.StatusInternalServerError, "Internal Server Error, Please Contact Administrator")

		if err := json.NewEncoder(w).Encode(responseError); err != nil {
			panic(err)
		}

	}
}

func ResponseServerError(code int, message string, w http.ResponseWriter, r *http.Request) {
	defer InternalServerError(w, r)

	w.Header().Set("Content-Type", "application/json")

	//TODO: Accept Http Error Code from callee function
	w.WriteHeader(http.StatusInternalServerError)

	responseError := NewResponseError(code, message)

	if err := json.NewEncoder(w).Encode(responseError); err != nil {
		panic(err)
	}
}
