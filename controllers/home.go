package controllers
import (
	"net/http"
	"html/template"
	"bridge/models/sessions"
	"bridge/models"
)
type Home struct {
	baseController
}
func(h Home) Get(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles(viewDir + "home.gtpl")
	cookie, _ := r.Cookie("username")
	session, err := SessionInstance.Get(sessions.Session{sessions.UserSession{models.User{}}}, cookie.Value)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	}else{
		t.Execute(w, session)
	}
}