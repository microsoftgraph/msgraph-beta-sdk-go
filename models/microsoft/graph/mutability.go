package graph
import (
    "strings"
    "errors"
)
type Mutability int

const (
    READWRITE_MUTABILITY Mutability = iota
    READONLY_MUTABILITY
    IMMUTABLE_MUTABILITY
    WRITEONLY_MUTABILITY
)

func (i Mutability) String() string {
    return []string{"READWRITE", "READONLY", "IMMUTABLE", "WRITEONLY"}[i]
}
func ParseMutability(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "READWRITE":
            return READWRITE_MUTABILITY, nil
        case "READONLY":
            return READONLY_MUTABILITY, nil
        case "IMMUTABLE":
            return IMMUTABLE_MUTABILITY, nil
        case "WRITEONLY":
            return WRITEONLY_MUTABILITY, nil
    }
    return 0, errors.New("Unknown Mutability value: " + v)
}
func SerializeMutability(values []Mutability) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
