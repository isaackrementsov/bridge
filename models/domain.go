package models
import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"bridge/utils"
	"errors"
)
type Domain struct {
	Name string
	Owner string
	Apps []App
}
func(d *Domain) GetByName(name string) error {
	connection, err := mgo.Dial("localhost:27017")
	defer connection.Close()
	utils.CheckErr("Error connecting to db: ", err)
	domains := connection.DB("bridge").C("domains")
	notFound := domains.Find(bson.M{"Name":name}).One(&d)
	return notFound
}
func(d *Domain) GetNewApp(a string) error {
	connection, err := mgo.Dial("localhost:27017")
	defer connection.Close()
	utils.CheckErr("Error connecting to db: ", err)
	apps := connection.DB("bridge").C("apps")
	domains := connection.DB("bridge").C("domains")
	err = apps.Update(bson.M{"Name":a}, bson.M{"$push":bson.M{"Domains":d.Name}})
	utils.CheckErr("Error updating app: ", err)
	app := App{}
	notFound := apps.Find(bson.M{"Name":a}).One(&app)
	err = domains.Update(bson.M{"$push":bson.M{"Apps":app}})
	utils.CheckErr("Error updating domain:", err)
	if notFound == nil {
		append(app, d.Apps)
	}
	return notFound
}
func(d *Domain) Save() error {
	connection, err := mgo.Dial("localhost:27017")
	defer connection.Close()
	utils.CheckErr("Error connecting to db: ", err)
	domains := connection.DB("bridge").C("domains")
	unique := isUnique("Name", d.Name, "domains", d)
	if unique {
		err = domains.Insert(d)
		utils := utils.CheckErr("Domain insert:", err)
		return nil
	}else{
		return errors.New("Domain name is not unique")
	}	
}
func(d *Domain) GetApps() ([]App, error){
	connection, err := mgo.Dial("localhost:27017")
	defer connection.Close()
	utils.CheckErr("Error connecting to db: ", err)
	apps := connection.DB("bridge").C("apps")
	var appArr []App
	notFound := apps.Find(bson.M{"domains":d}).All(&appArr)
	return appArr, notFound
}