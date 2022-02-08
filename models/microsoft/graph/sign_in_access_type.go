package graph
import (
    "strings"
    "errors"
)
// 
type SignInAccessType int

const (
    NONE_SIGNINACCESSTYPE SignInAccessType = iota
    B2BCOLLABORATION_SIGNINACCESSTYPE
    B2BDIRECTCONNECT_SIGNINACCESSTYPE
    MICROSOFTSUPPORT_SIGNINACCESSTYPE
    SERVICEPROVIDER_SIGNINACCESSTYPE
    UNKNOWNFUTUREVALUE_SIGNINACCESSTYPE
)

func (i SignInAccessType) String() string {
    return []string{"NONE", "B2BCOLLABORATION", "B2BDIRECTCONNECT", "MICROSOFTSUPPORT", "SERVICEPROVIDER", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseSignInAccessType(v string) (interface{}, error) {
    result := NONE_SIGNINACCESSTYPE
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_SIGNINACCESSTYPE
        case "B2BCOLLABORATION":
            result = B2BCOLLABORATION_SIGNINACCESSTYPE
        case "B2BDIRECTCONNECT":
            result = B2BDIRECTCONNECT_SIGNINACCESSTYPE
        case "MICROSOFTSUPPORT":
            result = MICROSOFTSUPPORT_SIGNINACCESSTYPE
        case "SERVICEPROVIDER":
            result = SERVICEPROVIDER_SIGNINACCESSTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_SIGNINACCESSTYPE
        default:
            return 0, errors.New("Unknown SignInAccessType value: " + v)
    }
    return &result, nil
}
func SerializeSignInAccessType(values []SignInAccessType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
