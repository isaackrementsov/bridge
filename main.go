package main
import (
	"net/http"
	"bridge/controllers"
	"bridge/models/sessions"
)
func main(){
	store := sessions.NewMemoryStore()
	controllers.InitializeSessions(store)
	http.Handle("/", controllers.Handler{controllers.Index{}})
	http.Handle("/signup", controllers.Handler{controllers.SignUp{}})
	http.Handle("/login", controllers.Handler{controllers.Login{}})
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("../src/bridge/assets"))))
	http.ListenAndServe(":8080", nil)
}