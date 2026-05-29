package server

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"vannucci.com/envtoschema/m/internal/infer"
)

func indexHandler(data PageData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data)
	}
}

func generateHandler(outputPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()

		schema := map[string]any{
			"$schema":    "http://json-schema.org/draft-04/schema#",
			"type":       "object",
			"properties": map[string]any{},
			"required":   []string{},
		}

		props := schema["properties"].(map[string]any)
		required := schema["required"].([]string)

		for _, key := range r.Form["key"] {
			t := r.FormValue("type_" + key)
			desc := r.FormValue("desc_" + key)
			isRequired := r.FormValue("required_"+key) == "true"

			prop := map[string]any{"type": t}
			if desc != "" {
				prop["description"] = desc
			}
			props[key] = prop

			if isRequired {
				required = append(required, key)
			}
		}

		schema["required"] = required

		out, err := json.MarshalIndent(schema, "", "  ")
		if err != nil {
			http.Error(w, "failed to marshal schema", http.StatusInternalServerError)
			return
		}

		if err := os.WriteFile(outputPath, out, 0644); err != nil {
			http.Error(w, "failed to write schema file", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`<p>Schema written to ` + outputPath + `.</p><script>setTimeout(() => { window.close() }, 1500)</script>`))

		go func() {
			time.Sleep(2 * time.Second)
			os.Exit(0)
		}()

	}

}

func ToFieldViews(elements []infer.InferredElement) []FieldView {
	views := make([]FieldView, 0, len(elements))
	for _, e := range elements {
		views = append(views, FieldView{
			Key:          e.Key,
			InferredType: infer.TypeToString(e.Candidate.Primary),
		})
	}
	return views
}
