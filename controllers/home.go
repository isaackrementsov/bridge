package controllers
import (
	"net/http"
	"html/template"
)
type Home struct {
	baseController
}
type PageData struct {
	Data map[string]interface{}
}
func(h Home) Get(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles(viewDir + "home.gtpl")
	cookie, _ := r.Cookie("username")
	session, err := SessionInstance.Get(cookie.Value)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	}else{
		t.Execute(w, session)
	}
}