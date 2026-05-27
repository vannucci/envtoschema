# Tomorrow's Todo

## 1. main.go — wire the pipeline

```go
path := argsWithoutProg[0]

data, err := read.ReadFile(path, 10<<20)
if err != nil { log.Fatal(err) }

err = validate.IsJSON(data)
if err != nil { log.Fatal(err) }

entries, err := infer.ParseFlat(data)
if err != nil { log.Fatal(err) }

elements := infer.Infer(entries)

server.Start(elements)
```

## 2. server.go — start the server

```go
func Start(elements []infer.InferredElement) {
    fields := toFieldViews(elements)
    data := PageData{Fields: fields}

    http.HandleFunc("GET /", indexHandler(data))
    http.HandleFunc("POST /generate", generateHandler)

    log.Println("listening on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func toFieldViews(elements []infer.InferredElement) []FieldView {
    views := make([]FieldView, 0, len(elements))
    for _, e := range elements {
        views = append(views, toFieldView(e))
    }
    return views
}
```

## 3. handlers.go — GET and POST

```go
func indexHandler(data PageData) http.HandlerFunc {
    tmpl := template.Must(template.ParseFiles("internal/server/templates/index.html"))
    return func(w http.ResponseWriter, r *http.Request) {
        tmpl.Execute(w, data)
    }
}

func generateHandler(w http.ResponseWriter, r *http.Request) {
    // parse form, build schema, return JSON
    // tomorrow's problem
}
```

## 4. index.html — select preselection

The template already handles this — just make sure
`InferredType` matches exactly: `"string"`, `"integer"`,
`"number"`, `"boolean"`. TypeToString handles that mapping.

## 5. toFieldView in handlers.go

```go
func toFieldView(e infer.InferredElement) FieldView {
    return FieldView{
        Key:          e.Key,
        InferredType: infer.TypeToString(e.Candidate.Primary),
    }
}
```

## Imports you'll need

```go
// server.go
"net/http"
"log"
"html/template"
"vannucci.com/envtoschema/m/internal/infer"

// main.go
"vannucci.com/envtoschema/m/internal/read"
"vannucci.com/envtoschema/m/internal/validate"
"vannucci.com/envtoschema/m/internal/infer"
"vannucci.com/envtoschema/m/internal/server"
```
