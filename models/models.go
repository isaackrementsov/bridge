package models
import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"reflect"
	"bridge/utils"
)
type Model struct { }
func(m *Model) IsUnique(val string, collection string) bool {
	r := reflect.ValueOf(m)
	field := reflect.Indirect(r).FieldByName(val)
	connection, err := mgo.Dial("localhost:27017")
	defer connection.Close()
	utils.CheckErr("Error connecting to db: ", err)
	c := connection.DB("bridge").C(collection)
	notFound := c.Find(bson.M{val:field}).One(&m)
	if notFound != nil {
		return false
	}else{
		return true
	}
}