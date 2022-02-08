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
    result := INORGANIZATION_ACCESSSCOPE
    switch strings.ToUpper(v) {
        case "INORGANIZATION":
            result = INORGANIZATION_ACCESSSCOPE
        case "NOTINORGANIZATION":
            result = NOTINORGANIZATION_ACCESSSCOPE
        default:
            return 0, errors.New("Unknown AccessScope value: " + v)
    }
    return &result, nil
}
func SerializeAccessScope(values []AccessScope) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
