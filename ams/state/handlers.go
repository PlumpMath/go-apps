// Copyright 2016 Jittakal Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package state

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/jittakal/go-apps/ams/response"
)

func Index(w http.ResponseWriter, r *http.Request) {
	defer response.InternalError(w)

	states, err := All()

	if err != nil {
		response.Error(http.StatusInternalServerError, err.Error(), w)
	} else {
		response.Ok(states, w)
	}
}

func IndexWithLink(w http.ResponseWriter, r *http.Request) {
	defer response.InternalError(w)

	states, err := All()

	if err != nil {
		response.Error(http.StatusInternalServerError, err.Error(), w)
	} else {
		state := states[0]
		responseState := ResponseState{
			State: state,
			Links: []HyperLink{
				HyperLink{
					Relation: "country",
					URL:      "/ams/country/v1/" + state.CountryId.Hex(),
					Method:   "GET",
				},
			},
		}

		response.Ok(responseState, w)
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	var state State
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &state); err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	if err := state.Create(); err != nil {
		fmt.Fprintf(w, "Failed to create new state"+err.Error())
	} else {
		fmt.Fprintf(w, "New State Created")
	}
}
