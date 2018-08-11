package models
import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"bridge/utils"
	"bridge/models/sessions"
	"errors"
)
type User struct {
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
func(u *User) Save() error {
	connection, err := mgo.Dial("localhost:27017")
	defer connection.Close()
	utils.CheckErr("Error connecting to db: ", err)	
	users := connection.DB("bridge").C("users")
	unique := isUnique("Username", u.Username, "users", u)
	if unique {
		err = users.Insert(&u)
		utils.CheckErr("User insert: ", err)
		return nil
	}else{
		return errors.New("Username is not unique")
	}
}
func(u *User) GetDomains() ([]Domain, error){
	connection, err := mgo.Dial("localhost:27017")
	defer connection.Close()
	utils.CheckErr("Error connecting to db: ", err)	
	domains := connection.DB("bridge").C("domains")
	var domainArr []Domain
	notFound := domains.Find(bson.M{"Owner":u.Username}).All(&domainArr)
	return domainArr, notFound
}
func(u *User) GetFromSession(s sessions.userSession){
	u.Username = s.Username
	u.Password = s.Password
	u.Email = s.Email
}