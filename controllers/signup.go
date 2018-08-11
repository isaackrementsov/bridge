package controllers
import (
	"net/http"
	"bridge/models"
	"strings"
)
type SignUpC struct { 
	baseController
}
func(s SignUpC) Get(w http.ResponseWriter, r *http.Request){
	render(w, "signup", nil)
}
func(s SignUpC) Post(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	user := models.User{strings.Join(r.Form["username"], ""), strings.Join(r.Form["password"], ""), strings.Join(r.Form["email"], "")}
	notUnique := user.Save()
	if notUnique != nil {
		http.Redirect(w, r, "/signup", 302)
	}else{
		ctr := LoginC{}
		ctr.Post(w, r)
	}
}