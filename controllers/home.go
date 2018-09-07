package controllers
import (
	"net/http"
	"bridge/models"
	"bridge/models/sessions"
)
type Home struct {
	baseController
}
type homePageData struct {
	Session sessions.UserSession
	DB []models.Domain
}
func(h Home) Get(w http.ResponseWriter, r *http.Request){
	session := checkSession(w, r)
	user := models.User{}
	user.GetFromSession(session)
	apps, _ := user.GetDomains()
	page := homePageData{session, apps}
	render(w, "home", page)
}