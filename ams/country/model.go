package country

import (
	"github.com/jittakal/go-apps/ams/common"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	host       = common.Host
	database   = common.Database
	collection = "country"
)

type Country struct {
	Id   bson.ObjectId `json:"id"  bson:"id,omitempty"`
	Code string        `json:"code"`
	Name string        `json:"name"`
}

func (c Country) Create() error {
	session, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	country := session.DB(database).C(collection)
	err = country.Insert(&c)
	return err
}

func All() ([]Country, error) {
	session, err := mgo.Dial("host")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	country := session.DB(database).C(collection)
	result := []Country{}
	err = country.Find(bson.M{}).All(&result)
	return result, err
}
