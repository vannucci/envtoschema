package server

import (
	"html/template"
	"log"
	"net/http"
	"os/exec"
	"time"
)

type FieldView struct {
	Key          string
	InferredType string
}

type PageData struct {
	Fields []FieldView
}

var tmpl = template.Must(template.ParseFiles("internal/server/templates/index.html"))

func Start(data PageData, outputPath string) {
	http.HandleFunc("GET /", indexHandler(data))
	http.HandleFunc("POST /generate", generateHandler(outputPath))
	log.Println("listening on :8080")
	go func() {
		time.Sleep(100 * time.Millisecond)
		exec.Command("open", "http://localhost:8080").Start() // macOS
	}()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
