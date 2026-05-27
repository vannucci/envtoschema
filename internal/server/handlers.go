package server

import (
	"vannucci.com/envtoschema/m/internal/infer"
)

func toFieldView(e infer.InferredElement) FieldView {
	return FieldView{
		Key:          e.Key,
		InferredType: TypeToString(e.Candidate.Primary),
	}
}
