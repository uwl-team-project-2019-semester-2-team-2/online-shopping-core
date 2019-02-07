package core

import (
	"../common"
	"net/http"
	"html/template"
	"github.com/go-chi/chi"
)

func Route(r chi.Router) {
	r.Get("/", handler)
	r.Route("/{productID}", func (r chi.Router) {
		}) 

}

func handler(w http.ResponseWriter, r *http.Request) {
	webpage := common.Page { "Core" }
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, webpage)
}
