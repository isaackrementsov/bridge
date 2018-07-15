package controllers
import (
	"net/http"
	"html/template"
	"gopkg.in/mgo.v2"
	"bridge/utils"
	"bridge/models"
	"strings"
)
type SignUp struct { }
func(S SignUp) ServeHTTP(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		S.get(w, r)
	}else if r.Method == "POST" {
		S.post(w, r)
	}
}
func(S SignUp) get(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("../src/bridge/views/signup.gtpl")
	t.Execute(w, nil)
}
func(S SignUp) post(w http.ResponseWriter, r *http.Request){
	server, err := mgo.Dial("localhost:27017")
	defer server.Close()
	utils.CheckErr(err)
	collection := server.DB("bridge").C("users")
	r.ParseForm()
	err = collection.Insert(&models.User{strings.Join(r.Form["username"], ""), strings.Join(r.Form["password"], ""), strings.Join(r.Form["email"], "")})
	utils.CheckErr(err)
	http.Redirect(w, r, "/login", 302)
}