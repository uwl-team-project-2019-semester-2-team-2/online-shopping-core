package marketing

import (
	"../common"
	"net/http"
	"html/template"
	"github.com/go-chi/chi"
)

func Route(r chi.Router) {
	r.Get("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	webpage := common.Page { "Marketing" }
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, webpage)
}
