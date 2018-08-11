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
	http.Handle("/signup", controllers.Handler{controllers.SignUpC{}})
	http.Handle("/login", controllers.Handler{controllers.LoginC{}})
	http.Handle("/home", controllers.Handler{controllers.Home{}})
	http.Handle("/apps", controllers.Handler{controllers.App{}})
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("../src/bridge/assets"))))
	http.ListenAndServe(":8080", nil)
}