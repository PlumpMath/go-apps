// Copyright 2016 Jittakal Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package main
package main

import (
	"log"
	"net/http"

	"github.com/jittakal/go-apps/ams/router"
)

func main() {
	router := router.NewRouter()
	log.Println("Starting server, access it with 'http://localhost:8080/{api-routes}'")
	log.Fatal(http.ListenAndServe(":8080", router))
}
