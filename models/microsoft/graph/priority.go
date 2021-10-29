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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_PRIORITY, nil
        case "HIGH":
            return HIGH_PRIORITY, nil
        case "LOW":
            return LOW_PRIORITY, nil
    }
    return 0, errors.New("Unknown Priority value: " + v)
}
func SerializePriority(values []Priority) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
