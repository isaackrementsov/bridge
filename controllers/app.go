package controllers
import (
	"net/http"
	"bridge/models/sessions"
	"bridge/models"
)
type App struct {
	baseController
}
type appPageData struct {
	Session sessions.UserSession
	DB []models.App
}
func(a App) Get(w http.ResponseWriter, r *http.Request){
	cookie, _ := r.Cookie("username")
	session, err := SessionInstance.Get(cookie.Value)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	}else{
		user := models.User{}
		user.GetFromSession(session)
		apps, _ := user.GetApps()
		page := appPageData{session, apps}
		render(w, "apps", page)
	}
}
func(a App) Patch(w http.ResponseWriter, r *http.Request){

}