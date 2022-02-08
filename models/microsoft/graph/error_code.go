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
    result := NOERROR_ERRORCODE
    switch strings.ToUpper(v) {
        case "NOERROR":
            result = NOERROR_ERRORCODE
        case "UNAUTHORIZED":
            result = UNAUTHORIZED_ERRORCODE
        case "NOTFOUND":
            result = NOTFOUND_ERRORCODE
        case "DELETED":
            result = DELETED_ERRORCODE
        default:
            return 0, errors.New("Unknown ErrorCode value: " + v)
    }
    return &result, nil
}
func SerializeErrorCode(values []ErrorCode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
