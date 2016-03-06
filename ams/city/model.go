package city

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	Host       = "localhost"
	Database   = "amstest"
	Collection = "city"
)

type City struct {
	Id   bson.ObjectId `json:"id"        bson:"_id,omitempty"`
	Code string        `json:"code"`
	Name string        `json:"name"`
}

func (c City) Create() error {
	session, err := mgo.Dial(Host)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	city := session.DB(Database).C("city")

	err = city.Insert(&c)
	return err
}

func (c City) Update() error {
	session, err := mgo.Dial(Host)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	city := session.DB(Database).C(Collection)

	err = city.Update(bson.M{"_id": c.Id}, c)
	return err
}

func FindById(id string) (City, error) {
	session, err := mgo.Dial(Host)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	city := session.DB(Database).C(Collection)
	result := City{}
	fmt.Println(id)

	err = city.FindId(bson.ObjectIdHex(id)).One(&result)
	return result, err
}

func FindByName(name string) ([]City, error) {
	session, err := mgo.Dial(Host)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	city := session.DB(Database).C("city")

	result := []City{}
	err = city.Find(bson.M{"name": name}).All(&result)
	return result, err
}

func All() ([]City, error) {
	session, err := mgo.Dial(Host)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	city := session.DB(Database).C("city")

	result := []City{}
	err = city.Find(bson.M{}).All(&result)
	return result, err
}
