package state

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/jittakal/go-apps/ams/common"
)

type HyperLink struct {
	Relation string `json:"rel"`
	URL      string `json:"url"`
	Method   string `json:"method"`
}

type ResponseState struct {
	State State       `json:"state"`
	Links []HyperLink `json:"links"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	defer common.InternalServerError(w, r)

	states, err := All()

	if err != nil {
		common.ResponseServerError(http.StatusInternalServerError, err.Error(), w, r)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err = json.NewEncoder(w).Encode(states); err != nil {
			panic(err)
		}
	}
}

func IndexWithLink(w http.ResponseWriter, r *http.Request) {
	defer common.InternalServerError(w, r)

	states, err := All()

	if err != nil {
		common.ResponseServerError(http.StatusInternalServerError, err.Error(), w, r)
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

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err = json.NewEncoder(w).Encode(responseState); err != nil {
			panic(err)
		}
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
