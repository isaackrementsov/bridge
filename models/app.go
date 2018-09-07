package models
import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"bridge/utils"
)
type App struct {
	Name string
	Description string
	Price int
	Domains []string
	Settings map[string]interface{}
}
func(a *App) GetByName(n string) error {
	connection, err := mgo.Dial("localhost:27017")
	defer connection.Close()
	utils.CheckErr("Error connecting to db: ", err)
	apps := connection.DB("bridge").C("apps")
	notFound := apps.Find(bson.M{"Name":n}).One(&a)
	return notFound
}
func GetApps() ([]App, error) {
	connection, err := mgo.Dial("localhost:27017")
	defer connection.Close()
	utils.CheckErr("Error connecting to db: ", err)
	apps := connection.DB("bridge").C("apps")
	var appArr []App
	notFound := apps.Find(bson.M{}).All(&appArr)
	return appArr, notFound
}