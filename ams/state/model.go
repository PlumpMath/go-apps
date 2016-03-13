package state

import (
	"github.com/jittakal/go-apps/ams/common"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	host       = common.Host
	database   = common.Database
	collection = "state"
)

type State struct {
	Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CountryId bson.ObjectId `json:"country_id"`
	Code      string        `json:"code"`
	Name      string        `json:"name"`
}

func (s State) Create() error {
	session, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	state := session.DB(database).C(collection)
	err = state.Insert(&s)
	return err
}

func All() ([]State, error) {
	session, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	state := session.DB(database).C(collection)
	result := []State{}
	err = state.Find(bson.M{}).All(&result)
	return result, err
}
