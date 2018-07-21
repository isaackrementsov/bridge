package controllers
import (
	"net/http"
	"strings"
	"html/template"
	"bridge/models"
	"bridge/models/sessions"
)
type Login struct { 
	baseController
}
func(l Login) Get(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles(viewDir + "login.gtpl")
	t.Execute(w, nil)
}
func(l Login) Post(w http.ResponseWriter, r *http.Request){
	user := models.User{}
	r.ParseForm()
	notFound := user.Login(strings.Join(r.Form["username"], ""), strings.Join(r.Form["password"], ""))
	var url string
	if notFound != nil {
		url = "/login"
	}else{
		url = "/home"
		userSession := sessions.Session{sessions.UserSession{user}}
		SessionInstance.Set(user.Username, userSession)
		uid := http.Cookie{Name:"username", Value:user.Username, Expires:getCookieExpiration()}
		http.SetCookie(w, &uid)
	} 
	http.Redirect(w, r, url, 302)
}