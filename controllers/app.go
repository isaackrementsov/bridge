package controllers
import (
	"net/http"
	"bridge/models"
)
type App struct {
	baseController
}
type appPageData struct {
	DB models.App
}
func(a App) Get(w http.ResponseWriter, r *http.Request){
	name := getParams(r, "/apps/")
	app := models.App{}
	app.GetByName(name)
	page := appPageData{app}
	render(w, "apps", page)
}
func(a App) Patch(w http.ResponseWriter, r *http.Request){
	
}