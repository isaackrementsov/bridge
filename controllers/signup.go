package controllers
import (
	"net/http"
	"html/template"
	"gopkg.in/mgo.v2"
	"bridge/utils"
	"bridge/models"
	"strings"
)
type SignUp struct { 
	baseController
}
func(s SignUp) get(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("../src/bridge/views/signup.gtpl")
	t.Execute(w, nil)
}
func(s SignUp) post(w http.ResponseWriter, r *http.Request){
	server, err := mgo.Dial("localhost:27017")
	defer server.Close()
	utils.CheckErr("Error dialing server:", err)
	users := server.DB("bridge").C("users")
	r.ParseForm()
	err = users.Insert(&models.User{strings.Join(r.Form["username"], ""), strings.Join(r.Form["password"], ""), strings.Join(r.Form["email"], "")})
	utils.CheckErr("User insert: ", err)
	http.Redirect(w, r, "/login", 302)
}