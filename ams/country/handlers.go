package country

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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

func Create(w http.ResponseWriter, r *http.Request) {
	var icountry Country
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &icountry); err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	if err := icountry.Create(); err != nil {
		fmt.Fprintf(w, "Failed to create new country"+err.Error())
	} else {
		fmt.Fprintf(w, "New Country Created")
	}
}
