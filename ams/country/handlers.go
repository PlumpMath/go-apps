// Copyright 2016 Jittakal Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package country

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jittakal/go-apps/ams/response"
)

func Index(w http.ResponseWriter, r *http.Request) {
	defer response.InternalError(w)

	countries, err := All()

	if err != nil {
		response.Error(http.StatusInternalServerError,
			err.Error(), w)
	} else {
		response.Ok(countries, w)
	}
}

func FindId(w http.ResponseWriter, r *http.Request) {
	defer response.InternalError(w)

	params := mux.Vars(r)

	id := params["id"]
	country, err := FindById(id)

	if err != nil {
		response.Error(http.StatusNotFound, err.Error(), w)
	} else {
		response.Ok(country, w)
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
