package controllers
import (
	"net/http"
	"html/template"
	"bridge/models/sessions"
)
type Home struct {
	baseController
}
func(h Home) Get(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles(viewDir + "home.gtpl")
	cookie, _ := r.Cookie("username")
	session, err := SessionInstance.Get(sessions.Session{sessions.UserSession{}}, cookie.Value)
	sessMap := sessions.GetMap(session)
	if err != nil || sessMap["Username"] == "" {
		http.Redirect(w, r, "/login", 302)
	}else{
		t.Execute(w, session)
	}
}