package schema

type Type int

const (
	TypeString Type = iota
	TypeInt
	TypeFloat32
	TypeBool
)

type TypeCandidate struct {
	Primary Type
	Alternate []Type
}

type Constraints struct {
	// numeric
	Min *float32
	Max *float32
	// string
	MaxLen *int
	Enum []string
}

type SchemaElement struct {
	Key string
	Candidate TypeCandidate
	FinalType *Type // what user confirmed
	Constraints Constraints // range, maxlen, enum, etc
	Required bool
}