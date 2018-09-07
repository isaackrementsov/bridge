package controllers
import (
	"bridge/models/sessions"
	"time"
	"net/http"
	"html/template"
	"strings"
)
var SessionInstance sessions.SessionStore
var viewDir string = "../src/bridge"
var CookieExpires time.Duration
type stdPageData struct { }
func SetViewDir(dir string){
	viewDir += dir + "/"
}
func getCookieExpiration() time.Time {
	return time.Now().Add(CookieExpires)
}
func render(w http.ResponseWriter, page string, data interface{}){
	t, _ := template.Must(template.ParseFiles(viewDir + "layout.gtpl", viewDir + page + ".gtpl"))
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
func checkSession(w http.ResponseWriter, r *http.Request) sessions.UserSession {
	session, err := getSession(r)
	if err != nil {
		defer http.Redirect(w, r, "/login", 302)
	}
	return session
}
func getParams(r *http.Request, url string) string {
	return strings.Split(r.URL.Path, url)[1]
}
func joinForm(r *http.Request) http.RequestForm {
	r.ParseForm()
	for key in r {
		r.Form[key] = strings.Join(r.Form[key], "")
	}
	return r.Form
}