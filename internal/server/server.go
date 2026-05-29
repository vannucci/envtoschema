package server

// import (
// 	"log"
// 	"net/http"
// )

// type FieldView struct {
// 	Key          string
// 	InferredType string
// }

// type PageData struct {
// 	Fields []FieldView
// }

// func Start(elements []InferredElement) {
// 	fields := toFieldViews(elements)
// 	data := PageData{Fields: fields}

// 	http.HandleFunc("GET /", indexHandler(data))
// 	http.HandleFunc("POST /generate", generateHandler)

// 	log.Println("listening on :8080")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// func toFieldViews(elements []InferredElement) []FieldView {
// 	views := make([]FieldView, 0, len(elements))
// 	for _, e := range elements {
// 		views = append(views, toFieldView(e))
// 	}
// 	return views
// }
