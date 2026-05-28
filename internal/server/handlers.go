package server

import (
	"net/http"
	"html/template"
	"vannucci.com/envtoschema/m/internal/infer"
)

func indexHandler(data PageData) http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles("internal/server/templates/index.html"))

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data)
	}
}

func generateHandler(w http.ResponseWriter, r *http.Request) {

}

func toFieldView(e InferredElement) FieldView {
	return FieldView{
		Key: e.Key,
		InferredType: infer.TypeToString(e.Candidate.Primary),
	}
}