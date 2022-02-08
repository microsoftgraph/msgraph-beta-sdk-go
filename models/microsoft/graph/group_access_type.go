package graph
import (
    "strings"
    "errors"
)
// 
type GroupAccessType int

const (
    NONE_GROUPACCESSTYPE GroupAccessType = iota
    PRIVATE_GROUPACCESSTYPE
    SECRET_GROUPACCESSTYPE
    PUBLIC_GROUPACCESSTYPE
)

func (i GroupAccessType) String() string {
    return []string{"NONE", "PRIVATE", "SECRET", "PUBLIC"}[i]
}
func ParseGroupAccessType(v string) (interface{}, error) {
    result := NONE_GROUPACCESSTYPE
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_GROUPACCESSTYPE
        case "PRIVATE":
            result = PRIVATE_GROUPACCESSTYPE
        case "SECRET":
            result = SECRET_GROUPACCESSTYPE
        case "PUBLIC":
            result = PUBLIC_GROUPACCESSTYPE
        default:
            return 0, errors.New("Unknown GroupAccessType value: " + v)
    }
    return &result, nil
}
func SerializeGroupAccessType(values []GroupAccessType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
