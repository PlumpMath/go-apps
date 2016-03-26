// Copyright 2016 Jittakal Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package city

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	cities, err := All()

	if err != nil {
		fmt.Fprintf(w, err.Error())
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err = json.NewEncoder(w).Encode(cities); err != nil {
			panic(err)
		}
	}
}

func ById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	city, err := FindById(id)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(city); err != nil {
		panic(err)
	}
}

func ByName(w http.ResponseWriter, r *http.Request) {
	var icity City
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &icity); err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	cities, err := FindByName(icity.Name)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(cities); err != nil {
		panic(err)
	}

}

func Create(w http.ResponseWriter, r *http.Request) {
	var icity City
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &icity); err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	if err := icity.Create(); err != nil {
		fmt.Fprintf(w, "Failed to create new city"+err.Error())
	} else {
		fmt.Fprintf(w, "New City Created")
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	var icity City
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &icity); err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	if err := icity.Update(); err != nil {
		fmt.Fprintf(w, "Failed to create new city"+err.Error())
	} else {
		fmt.Fprintf(w, "City Updated")
	}
}
