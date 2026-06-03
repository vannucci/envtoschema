package schema

import "testing"

func TestSchemaBuilder(t *testing.T) {
	t.Run("empty schema", func(t *testing.T) {
		schemaBuilder := New()
		schema := schemaBuilder.Build()

		if schema.Type != "object" {
			t.Fatalf("schema type is not object %v", schema.Type)
		}
	})

	t.Run("scalar property appears in schema", func(t *testing.T) {
		b := New()
		b.WriteScalar("host", TypeString)
		schema := b.Build()

		if _, ok := schema.Properties["host"]; !ok {
			t.Fatal("expected 'host' in properties")
		}
	})

}
