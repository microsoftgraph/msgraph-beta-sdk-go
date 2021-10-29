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
    switch strings.ToUpper(v) {
        case "DEFAULT":
            return DEFAULT_CONTENTFORMAT, nil
        case "EMAIL":
            return EMAIL_CONTENTFORMAT, nil
    }
    return 0, errors.New("Unknown ContentFormat value: " + v)
}
func SerializeContentFormat(values []ContentFormat) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
