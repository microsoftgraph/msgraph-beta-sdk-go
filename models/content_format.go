package models
import (
    "errors"
)
// 
type ContentFormat int

const (
    DEFAULTESCAPED_CONTENTFORMAT ContentFormat = iota
    EMAIL_CONTENTFORMAT
)

func (i ContentFormat) String() string {
    return []string{"default", "email"}[i]
}
func ParseContentFormat(v string) (any, error) {
    result := DEFAULTESCAPED_CONTENTFORMAT
    switch v {
        case "default":
            result = DEFAULTESCAPED_CONTENTFORMAT
        case "email":
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
func (i ContentFormat) isMultiValue() bool {
    return false
}
