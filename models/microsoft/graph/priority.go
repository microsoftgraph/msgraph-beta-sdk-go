package graph
import (
    "strings"
    "errors"
)
// 
type Priority int

const (
    NONE_PRIORITY Priority = iota
    HIGH_PRIORITY
    LOW_PRIORITY
)

func (i Priority) String() string {
    return []string{"NONE", "HIGH", "LOW"}[i]
}
func ParsePriority(v string) (interface{}, error) {
    result := NONE_PRIORITY
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_PRIORITY
        case "HIGH":
            result = HIGH_PRIORITY
        case "LOW":
            result = LOW_PRIORITY
        default:
            return 0, errors.New("Unknown Priority value: " + v)
    }
    return &result, nil
}
func SerializePriority(values []Priority) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
