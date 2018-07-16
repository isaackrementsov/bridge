package controllers
import (
	"net/http"
	"html/template"
)
type Index struct { 
	baseController
}
func(i Index) Get(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("../src/bridge/views/index.gtpl")
	t.Execute(w, nil)
}
