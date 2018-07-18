package controllers
import (
	"net/http"
	"html/template"
	"bridge/models"
	"strings"
)
type SignUp struct { 
	baseController
}
func(s SignUp) Get(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles(viewDir + "signup.gtpl")
	t.Execute(w, nil)
}
func(s SignUp) Post(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	user := models.User{models.Model{}, strings.Join(r.Form["username"], ""), strings.Join(r.Form["password"], ""), strings.Join(r.Form["email"], "")}
	notUnique := user.SignUp()
	if notUnique != nil {
		http.Redirect(w, r, "/signup", 302)
	}else{
		Index{}.Post(w, r)
	}
}