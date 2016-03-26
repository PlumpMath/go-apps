// Copyright 2016 Jittakal Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package country

import (
	"errors"
	"fmt"

	"github.com/jittakal/go-apps/ams/common"
	"gopkg.in/mgo.v2/bson"
)

const (
	collection = "country"
)

type Country struct {
	Id   bson.ObjectId `json:"id"  bson:"_id,omitempty"`
	Code string        `json:"code"`
	Name string        `json:"name"`
}

func (c Country) Create() error {
	session, database := common.Session()
	defer common.Close(session)

	country := database.C(collection)
	return country.Insert(&c)
}

func (c Country) Update() error {
	session, database := common.Session()
	defer common.Close(session)

	country := database.C(collection)
	return country.Update(bson.M{"_id": c.Id}, c)
}

func FindById(id string) (Country, error) {
	session, database := common.Session()
	defer common.Close(session)

	country := database.C(collection)
	result := Country{}

	if bson.IsObjectIdHex(id) {
		err := country.FindId(bson.ObjectIdHex(id)).One(&result)
		return result, err
	} else {
		return result, errors.New(fmt.Sprintf("Invalid Object Id: %s", id))
	}
}

func All() ([]Country, error) {
	session, database := common.Session()
	defer common.Close(session)

	country := database.C(collection)
	result := []Country{}
	if err := country.Find(bson.M{}).All(&result); err != nil {
		return nil, err
	}

	return result, nil
}
