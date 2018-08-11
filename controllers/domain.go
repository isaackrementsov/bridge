package controllers
import (
	"net/http"
	"bridge/models/sessions"
	"bridge/models"
	"errors"
)
type Domain struct {
	baseController
}
type domainPageData struct {
	Session sessions.UserSession
	DB models.Domain
}
func(d Domain) Get(w http.ResponseWriter, r *http.Request){
	session, err := getSession(r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	}else{
		name := split(r.URL.Path, "/domains/")[1]
		domain := models.Domain{}
		_ := domain.GetByName(name)
		page := domainPageData{session, domain}
		render(w, "domain", page)
	}
}
func(d Domain) Post(w http.ResponseWriter, r *http.Request){
	session, err := getSession(r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	}else{
		id := split(r.URL.Path, "/domains/")
		r.ParseForm()
		res, reqErr := http.Get("http://" + r.Form.Name + "/" + id)
		res.ParseForm()
		if reqErr == nil || res.Body == r.Form.Random {
			r.ParseForm()
			domain := models.Domain{r.Form.Name, session.Username, nil}
			err = Domain.Save()
		}else{

		}
	}
}
func(d Domain) Patch(w http.ResponseWriter, r *http.Request){
	
}
func(d Domain) Delete(w http.ResponseWriter, r *http.Request){
	
}