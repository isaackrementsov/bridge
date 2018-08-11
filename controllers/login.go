package controllers
import (
	"net/http"
	"strings"
	"bridge/models"
	"bridge/models/sessions"
)
type LoginC struct { 
	baseController
}
func(l LoginC) Get(w http.ResponseWriter, r *http.Request){
	render(w, "login", nil)
}
func(l LoginC) Post(w http.ResponseWriter, r *http.Request){
	user := models.User{}
	r.ParseForm()
	notFound := user.Login(strings.Join(r.Form["username"], ""), strings.Join(r.Form["password"], ""))
	var url string
	if notFound != nil {
		url = "/login"
	}else{
		url = "/home"
		s := sessions.UserSession{user.Username, user.Email, user.Password, nil}
		makeSession(w, s)
	} 
	http.Redirect(w, r, url, 302)
}