package country

import (
	"encoding/json"
	"net/http"
)

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()

	countries, err := All()

	if err != nil {
		//		fmt.Fprintf(w, err.Error())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)

		responseError := ResponseError{
			http.StatusInternalServerError,
			err.Error(),
		}
		if err = json.NewEncoder(w).Encode(responseError); err != nil {
			panic(err)
		}

	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err = json.NewEncoder(w).Encode(countries); err != nil {
			panic(err)
		}
	}
}
