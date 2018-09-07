package controllers
import (
	"net/http"
	"bridge/models/sessions"
	"bridge/models"
	"strings"
)
type Domain struct {
	baseController
}
type domainPageData struct {
	Session sessions.UserSession
	DB models.Domain
}
func(d Domain) Get(w http.ResponseWriter, r *http.Request){
	session := checkSession(w, r)
	name := getParams(r, "/domains/")
	domain := models.Domain{}
	domain.GetByName(name)
	page := domainPageData{session, domain}
	render(w, "domain", page)
}
func(d Domain) Post(w http.ResponseWriter, r *http.Request){
	/*session := checkSession(w, r)
	r.ParseForm()
	id := strings.Join(r.Form["Random"], "")
	res, reqErr := http.Get("http://" + strings.Join(r.Form["Name"], "") + "/" + id)
	res.ParseForm()
	if reqErr == nil || res.Body == r.Form.Random {
		r.ParseForm()
		domain := models.Domain{r.Form.Name, session.Username, nil}
		err = domain.Save()
	}else{
		http.Redirect(w, r, "/domains/", 302)
	}*/
}
func(d Domain) Patch(w http.ResponseWriter, r *http.Request){
	session := checkSession(w, r)
	r.ParseForm()
	domainName := strings.Join(r.Form["Domain"], "")
	appName := strings.Join(r.Form["App"], "")
	domain := models.Domain{}
	notFound := domain.Get(domainName, session.Username)
	if notFound == nil {
		domain.GetNewApp(appName)
		http.Redirect(w, r, "/domain/" + domainName, 302)
	}else{
		http.Redirect(w, r, "/app/" + appName, 302)
	}
}
func(d Domain) Delete(w http.ResponseWriter, r *http.Request){
	
}