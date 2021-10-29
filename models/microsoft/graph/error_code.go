package graph
import (
    "strings"
    "errors"
)
// 
type ErrorCode int

const (
    NOERROR_ERRORCODE ErrorCode = iota
    UNAUTHORIZED_ERRORCODE
    NOTFOUND_ERRORCODE
    DELETED_ERRORCODE
)

func (i ErrorCode) String() string {
    return []string{"NOERROR", "UNAUTHORIZED", "NOTFOUND", "DELETED"}[i]
}
func ParseErrorCode(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NOERROR":
            return NOERROR_ERRORCODE, nil
        case "UNAUTHORIZED":
            return UNAUTHORIZED_ERRORCODE, nil
        case "NOTFOUND":
            return NOTFOUND_ERRORCODE, nil
        case "DELETED":
            return DELETED_ERRORCODE, nil
    }
    return 0, errors.New("Unknown ErrorCode value: " + v)
}
func SerializeErrorCode(values []ErrorCode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
