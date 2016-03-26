// Copyright 2016 Jittakal Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package common

import "gopkg.in/mgo.v2"

const (
	Host     = "localhost"
	Database = "amstest"
)

func Session() (*mgo.Session, *mgo.Database) {
	session, err := mgo.Dial(Host)
	if err != nil {
		panic(err)
	}
	return session, session.DB(Database)
}

func Close(session *mgo.Session) {
	if session != nil {
		defer session.Close()
	}
}
