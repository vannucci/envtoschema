package server

type FieldView struct {
	Key          string
	InferredType string
}

type PageData struct {
	Fields []FieldView
}
