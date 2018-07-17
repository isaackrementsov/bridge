package main
import (
	"net/http"
	"bridge/controllers"
	"bridge/models/sessions"
	"time"
)
func main(){
	store := sessions.NewMemoryStore()
	controllers.SessionInstance = store
	controllers.CookieExpires = 365 * 24 * time.Hour
	controllers.SetViewDir("/views")
	http.Handle("/", controllers.Handler{controllers.Index{}})
	http.Handle("/signup", controllers.Handler{controllers.SignUp{}})
	http.Handle("/login", controllers.Handler{controllers.Login{}})
	http.Handle("/home", controllers.Handler{controllers.Home{}})
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("../src/bridge/assets"))))
	http.ListenAndServe(":8080", nil)
}