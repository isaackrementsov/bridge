package models
import (
	"bridge/utils"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)
func isUnique(key string, val string, coll string, m interface{}) bool {
	connection, err := mgo.Dial("localhost:27017")
	defer connection.Close()
	utils.CheckErr("Error connecting to db: ", err)
	collection := connection.DB("bridge").C(coll)
	match := collection.Find(bson.M{key:val}).One(&m)
	if match == nil {
		return false
	}else{
		return true
	}
}