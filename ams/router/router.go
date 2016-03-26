// Copyright 2016 Jittakal Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package router

import (
	"github.com/gorilla/mux"
	"github.com/jittakal/go-apps/ams/city"
	"github.com/jittakal/go-apps/ams/country"
	"github.com/jittakal/go-apps/ams/state"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range country.Routes {
		router.PathPrefix("/ams/country").
			Methods(route.Method).
			Path(route.Pattern).
			Handler(route.HandlerFunc)
	}

	for _, route := range state.Routes {
		router.PathPrefix("/ams/state").
			Methods(route.Method).
			Path(route.Pattern).
			Handler(route.HandlerFunc)
	}

	for _, route := range city.Routes {
		router.PathPrefix("/ams/city").
			Methods(route.Method).
			Path(route.Pattern).
			Handler(route.HandlerFunc)
	}

	return router
}
