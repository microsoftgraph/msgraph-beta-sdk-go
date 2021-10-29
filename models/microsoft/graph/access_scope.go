package graph
import (
    "strings"
    "errors"
)
// 
type AccessScope int

const (
    INORGANIZATION_ACCESSSCOPE AccessScope = iota
    NOTINORGANIZATION_ACCESSSCOPE
)

func (i AccessScope) String() string {
    return []string{"INORGANIZATION", "NOTINORGANIZATION"}[i]
}
func ParseAccessScope(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "INORGANIZATION":
            return INORGANIZATION_ACCESSSCOPE, nil
        case "NOTINORGANIZATION":
            return NOTINORGANIZATION_ACCESSSCOPE, nil
    }
    return 0, errors.New("Unknown AccessScope value: " + v)
}
func SerializeAccessScope(values []AccessScope) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
