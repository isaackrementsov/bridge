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
func(d *Domain) Get(n string, o string) error {
	connection, err := mgo.Dial("localhost:27017")
	defer connection.Close()
	utils.CheckErr("Error connecting to db: ", err)
	domains := connection.DB("bridge").C("domains")
	notFound := domains.Find(bson.M{"Name":n, "Owner":o}).One(&d)
	return notFound
}
func(d *Domain) GetNewApp(a string) error {
	connection, err := mgo.Dial("localhost:27017")
	defer connection.Close()
	utils.CheckErr("Error connecting to db: ", err)
	domains := connection.DB("bridge").C("domains")
	err = apps.Update(bson.M{"Name":a}, bson.M{"$push":bson.M{"Domains":d.Name}})
	utils.CheckErr("Error updating app: ", err)
	app := App{}
	notFound := app.GetByName(a)
	err = domains.Update(bson.M{"Name":d.Name, "Owner":d.Owner}, bson.M{"$push":bson.M{"Apps":app}})
	utils.CheckErr("Error updating domain:", err)
	if notFound == nil {
		d.Apps = append(d.Apps, app)
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
		utils.CheckErr("Domain insert:", err)
		return nil
	}else{
		return errors.New("Domain name is not unique")
	}	
}
func(d *Domain) ChangeAppSettings(n string, m map[string]interface{}){
	connection, err := mgo.Dial("localhost:27017")
	defer connection.Close()
	utils.CheckErr("Error connecting to db: ", err)
	domains := connection.DB("bridge").C("domains")
	app := d.Apps[Name == n]
	for key in m {
		if app.Settings[key] != nil {
			app.Settings[key] = m[key]
		}
	}
	err = domains.Update(bson.M{"Name":d.Name, "Owner":d.Owner}, json.Unmarshal(app.(bson.M)))
	utils.CheckErr("Error updating domain: ", err)
}