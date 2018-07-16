package controllers
import (
	"net/http"
	"strings"
	"html/template"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"bridge/utils"
	"bridge/models"
)
type Login struct { 
	baseController
}
func(l Login) Get(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("../src/bridge/views/login.gtpl")
	t.Execute(w, nil)
}
func(l Login) Post(w http.ResponseWriter, r *http.Request){
	server, err := mgo.Dial("localhost:27017")
	defer server.Close()
	utils.CheckErr("Error dialing server:", err)
	users := server.DB("bridge").C("users")
	user := models.User{}
	r.ParseForm()
	notFound := users.Find(bson.M{"username":strings.Join(r.Form["username"], ""), "password":strings.Join(r.Form["password"], "")}).One(&user)
	url := "/home"
	if notFound != nil {
		url = "/login"
	}
	http.Redirect(w, r, url, 302)
}