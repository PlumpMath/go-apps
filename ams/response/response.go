package response

import (
	"encoding/json"
	"net/http"
)

const (
	ContentType     = "Content-Type"
	JsonContentType = "application/json;charset=UTF-8"
)

type errorResponse struct {
	Error errorDetail `json:"error"`
}

type errorDetail struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func newErrorResponse(code int, message string) errorResponse {
	return errorResponse{
		Error: errorDetail{
			Code:    code,
			Message: message,
		},
	}
}

func encode(v interface{}, w http.ResponseWriter) {
	if err := json.NewEncoder(w).Encode(v); err != nil {
		panic(err)
	}
}

func InternalError(w http.ResponseWriter) {
	if r := recover(); r != nil {
		w.Header().Set(ContentType, JsonContentType)
		w.WriteHeader(http.StatusInternalServerError)
		errorResponse := newErrorResponse(http.StatusInternalServerError, "Internal Server Error, Please Contact Administrator")
		encode(errorResponse, w)
	}
}

func Error(statusCode int, message string, w http.ResponseWriter) {
	defer InternalError(w)

	w.Header().Set(ContentType, JsonContentType)
	w.WriteHeader(statusCode)

	errorResponse := newErrorResponse(statusCode, message)
	encode(errorResponse, w)
}

func Ok(v interface{}, w http.ResponseWriter) {
	defer InternalError(w)

	w.Header().Set(ContentType, JsonContentType)
	w.WriteHeader(http.StatusOK)
	encode(v, w)
}
