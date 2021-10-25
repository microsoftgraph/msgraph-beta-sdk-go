package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_GROUPACCESSTYPE, nil
        case "PRIVATE":
            return PRIVATE_GROUPACCESSTYPE, nil
        case "SECRET":
            return SECRET_GROUPACCESSTYPE, nil
        case "PUBLIC":
            return PUBLIC_GROUPACCESSTYPE, nil
    }
    return 0, errors.New("Unknown GroupAccessType value: " + v)
}
func SerializeGroupAccessType(values []GroupAccessType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
