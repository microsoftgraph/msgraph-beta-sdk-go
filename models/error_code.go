package models
import (
    "errors"
)
// 
type ErrorCode int

const (
    // Default Value to indicate no error.
    NOERROR_ERRORCODE ErrorCode = iota
    // The current user does not have access due to lack of RBAC permissions on the resource.
    UNAUTHORIZED_ERRORCODE
    // The current user does not have access due to lack of RBAC Scope Tags on the resource.
    NOTFOUND_ERRORCODE
    // The resource has been deleted.
    DELETED_ERRORCODE
)

func (i ErrorCode) String() string {
    return []string{"noError", "unauthorized", "notFound", "deleted"}[i]
}
func ParseErrorCode(v string) (any, error) {
    result := NOERROR_ERRORCODE
    switch v {
        case "noError":
            result = NOERROR_ERRORCODE
        case "unauthorized":
            result = UNAUTHORIZED_ERRORCODE
        case "notFound":
            result = NOTFOUND_ERRORCODE
        case "deleted":
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
func (i ErrorCode) isMultiValue() bool {
    return false
}
