package infer

// InferType examines a raw string value from the env file and returns the most
// specific type candidate possible

// Preconditions:
// 		- input is a non-empty string (empty strings handled upstream)

// Postconditions:
// 		- returns exactly one TypeCandidate
//		- never returns an error; ambiguity is encoded in the candidate itself

// Ambiguity rules:
// 		- "true" / "false" (case-insensitive) -> TypeAmbiguous{Bool, String}
// 		- valid integer string -> TypeAmbiguous{Int, Float32}
//		- valid float string -> TypeFloat32
// 		- anything else -> TypeString

// Invariants:
// 		- same input always produces same output (pure, no side effects)
//		- caller is responsible for resolving TypeAmbiguous via user input

func InferType(value string) TypeCandidate