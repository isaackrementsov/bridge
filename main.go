package main
import (
	"net/http"
	"bridge/controllers"
	"bridge/models"
)
func main(){
	session := models.Store
	http.Handle("/", controllers.Handler{controllers.Index{}})
	http.Handle("/signup", controllers.Handler{controllers.SignUp{session, controllers.BaseController{}}})
	http.Handle("/login", controllers.Handler{controllers.Login{session, controllers.BaseController{}}})
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("../src/bridge/assets"))))
	http.ListenAndServe(":8080", nil)
}