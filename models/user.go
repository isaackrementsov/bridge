package models
import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"bridge/utils"
)
type User struct {
	Username string
	Password string
	Email string
}
func(u User) Login(username string, password string) error {
	connection, err := mgo.Dial("localhost:27017")
	defer connection.Close()
	utils.CheckErr("Error connecting to db: ", err)
	users := connection.DB("bridge").C("users")
	notFound := users.Find(bson.M{"username":username, "password":password}).One(&u)
	return notFound
}
func(u User) SignUp(){
	connection, err := mgo.Dial("localhost:27017")
	defer connection.Close()
	utils.CheckErr("Error connecting to db: ", err)	
	users := connection.DB("bridge").C("users")
	err = users.Insert(&u)
	utils.CheckErr("User insert: ", err)
}