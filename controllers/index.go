package controllers
import "net/http"
type Index struct { 
	baseController
}
func(i Index) Get(w http.ResponseWriter, r *http.Request){
	render(w, "index", nil)
}
