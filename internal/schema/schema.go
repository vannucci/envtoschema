package schema

import (
	"encoding/json"
	"os"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

type SchemaNode struct {
	Type       string
	Properties map[string]*SchemaNode
	Items      *SchemaNode
	Required   []string
}

type SchemaBuilder struct {
	root    *SchemaNode
	current *SchemaNode
}

type ScalarType string

const (
	TypeString ScalarType = "string"
	TypeInt    ScalarType = "integer"
	TypeFloat  ScalarType = "number"
	TypeBool   ScalarType = "boolean"
)

type Schema struct {
	SchemaType string         `json:"$schema"`
	Type       string         `json:"type"`
	Properties map[string]any `json:"properties"`
	Required   []string       `json:"required"`
}

func New() *SchemaBuilder {
	root := &SchemaNode{
		Type:       "object",
		Properties: make(map[string]*SchemaNode),
		Required:   []string{},
	}

	return &SchemaBuilder{root: root, current: root}
}

func (s *SchemaBuilder) WriteScalar(key string, typ ScalarType) {
	s.current.Properties[key] = &SchemaNode{Type: string(typ)}
}

func (s *SchemaBuilder) WriteArrayItems(typ ScalarType) {
	s.current.Items = &SchemaNode{Type: string(typ)}
}

func (s *SchemaBuilder) WriteObject(key string) *SchemaBuilder {
	newNode := &SchemaNode{
		Type:       "object",
		Properties: make(map[string]*SchemaNode),
	}

	s.current.Properties[key] = newNode

	newBuilder := &SchemaBuilder{
		root:    s.root,
		current: newNode,
	}

	return newBuilder
}

func (s *SchemaBuilder) WriteArray(key string) *SchemaBuilder {
	newNode := &SchemaNode{
		Type: "array",
	}

	s.current.Properties[key] = newNode

	newBuilder := &SchemaBuilder{
		root:    s.root,
		current: newNode,
	}

	return newBuilder

}

func (s *SchemaBuilder) Build() Schema {
	schema := Schema{
		SchemaType: "http://json-schema.org/draft-04/schema#",
		Type:       s.root.Type,
		Properties: make(map[string]any),
		Required:   []string{},
	}

	for k, v := range s.root.Properties {
		schema.Properties[k] = buildNode(v)
	}

	return schema

}

func buildNode(node *SchemaNode) map[string]any {
	out := map[string]any{"type": node.Type}

	switch node.Type {
	case "object":
		props := map[string]any{}
		for k, v := range node.Properties {
			props[k] = buildNode(v)
		}
	case "array":
		if node.Items != nil {
			out["items"] = buildNode(node.Items)
		}
	}

	return out

}

func ValidateConfig(config string, schema string) error {
	compiler := jsonschema.NewCompiler()

	sch, err := compiler.Compile(schema)
	if err != nil {
		return err
	}

	c, err := os.Open(config)
	if err != nil {
		return err
	}
	defer c.Close()

	var instance any
	if err := json.NewDecoder(c).Decode(&instance); err != nil {
		return err
	}

	if err = sch.Validate(instance); err != nil {
		return err
	}

	return nil

}
