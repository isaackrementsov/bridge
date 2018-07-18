package models
import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"bridge/utils"
	"errors"
)
type User struct {
	Model
	Username string
	Password string
	Email string
}
func(u *User) Login(username string, password string) error {
	connection, err := mgo.Dial("localhost:27017")
	defer connection.Close()
	utils.CheckErr("Error connecting to db: ", err)
	users := connection.DB("bridge").C("users")
	notFound := users.Find(bson.M{"username":username, "password":password}).One(&u)
	return notFound
}
func(u *User) SignUp() error {
	connection, err := mgo.Dial("localhost:27017")
	defer connection.Close()
	utils.CheckErr("Error connecting to db: ", err)	
	users := connection.DB("bridge").C("users")
	unique := u.IsUnique("Name", "users")
	if unique {
		err = users.Insert(&u)
		utils.CheckErr("User insert: ", err)
		return nil
	}else{
		return errors.New("Username is not unique")
	}
}