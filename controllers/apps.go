package controllers
import (
	"bridge/models"
	"bridge/models/sessions"
	"net/http"
)
type Apps struct {
	baseController
}
type appsPageData struct {
	Session sessions.UserSession
	DB []models.App
}
func(a Apps) Get(w http.ResponseWriter, r *http.Request){
	//session := checkSession(w, r)
	//apps, _ := models.GetApps()
}