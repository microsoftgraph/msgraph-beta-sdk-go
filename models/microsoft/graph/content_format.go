package graph
import (
    "strings"
    "errors"
)
// 
type ContentFormat int

const (
    DEFAULT_CONTENTFORMAT ContentFormat = iota
    EMAIL_CONTENTFORMAT
)

func (i ContentFormat) String() string {
    return []string{"DEFAULT", "EMAIL"}[i]
}
func ParseContentFormat(v string) (interface{}, error) {
    result := DEFAULT_CONTENTFORMAT
    switch strings.ToUpper(v) {
        case "DEFAULT":
            result = DEFAULT_CONTENTFORMAT
        case "EMAIL":
            result = EMAIL_CONTENTFORMAT
        default:
            return 0, errors.New("Unknown ContentFormat value: " + v)
    }
    return &result, nil
}
func SerializeContentFormat(values []ContentFormat) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
