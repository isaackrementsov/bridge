package main
import (
	"net/http"
	"bridge/controllers"
)
func main(){
	http.Handle("/", controllers.Index{})
	http.Handle("/signup", controllers.SignUp{})
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("../src/bridge/assets"))))
	http.ListenAndServe(":8080", nil)
}