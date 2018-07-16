package controllers
import "net/http"
type Handler struct {
	Controller
}
func(h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	switch r.Method {
		case "GET":
			h.Get(w, r)
		case "POST":
			h.Post(w, r)
		case "PUT":
			h.Put(w, r)
		case "PATCH":
			h.Patch(w, r)
		case "DELETE":
			h.Delete(w, r)
	}
}
type Controller interface {
	Get(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
	Put(w http.ResponseWriter, r *http.Request)
	Patch(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}