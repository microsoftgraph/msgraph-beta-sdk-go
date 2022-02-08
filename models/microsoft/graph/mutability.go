package graph
import (
    "strings"
    "errors"
)
// 
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
    result := READWRITE_MUTABILITY
    switch strings.ToUpper(v) {
        case "READWRITE":
            result = READWRITE_MUTABILITY
        case "READONLY":
            result = READONLY_MUTABILITY
        case "IMMUTABLE":
            result = IMMUTABLE_MUTABILITY
        case "WRITEONLY":
            result = WRITEONLY_MUTABILITY
        default:
            return 0, errors.New("Unknown Mutability value: " + v)
    }
    return &result, nil
}
func SerializeMutability(values []Mutability) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
