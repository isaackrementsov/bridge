package controllers
import (
	"bridge/models/sessions"
	"time"
	"net/http"
	"html/template"
)
var SessionInstance sessions.SessionStore
var viewDir string = "../src/bridge"
var CookieExpires time.Duration
func SetViewDir(dir string){
	viewDir += dir + "/"
}
func getCookieExpiration() time.Time {
	return time.Now().Add(CookieExpires)
}
func render(w http.ResponseWriter, page string, data interface{}){
	t, _ := template.ParseFiles(viewDir + page + ".gtpl")
	t.Execute(w, &data)
}
func getSession(r *http.Request) (sessions.UserSession, error){
	cookie, _ := r.Cookie("username")
	session, err := SessionInstance.Get(cookie.Value)
	return session, err
}
func makeSession(w http.ResponseWriter, u sessions.UserSession){
	uid := http.Cookie{Name:"username", Value:u.Username, Expires:getCookieExpiration()}
	SessionInstance.Set(u, u.Username)
	http.SetCookie(w, &uid)
}
func updateSession(w http.ResponseWriter, u session.UserSession){
	
}