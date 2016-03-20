// Copyright 2016 @jittakal
package country

import (
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
