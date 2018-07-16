package controllers
import "net/http"
type baseController struct { }
func(b baseController) notAllowed(w http.ResponseWriter){
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("405- Method Not Allowed"))
}
func(b baseController) Get(w http.ResponseWriter, r *http.Request){
	b.notAllowed(w)
}
func(b baseController) Post(w http.ResponseWriter, r *http.Request){
	b.notAllowed(w)
}
func(b baseController) Put(w http.ResponseWriter, r *http.Request){
	b.notAllowed(w)
}
func(b baseController) Patch(w http.ResponseWriter, r *http.Request){
	b.notAllowed(w)
}
func(b baseController) Delete(w http.ResponseWriter, r *http.Request){
	b.notAllowed(w)
}
