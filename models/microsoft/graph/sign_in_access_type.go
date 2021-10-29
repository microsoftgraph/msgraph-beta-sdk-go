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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_SIGNINACCESSTYPE, nil
        case "B2BCOLLABORATION":
            return B2BCOLLABORATION_SIGNINACCESSTYPE, nil
        case "B2BDIRECTCONNECT":
            return B2BDIRECTCONNECT_SIGNINACCESSTYPE, nil
        case "MICROSOFTSUPPORT":
            return MICROSOFTSUPPORT_SIGNINACCESSTYPE, nil
        case "SERVICEPROVIDER":
            return SERVICEPROVIDER_SIGNINACCESSTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_SIGNINACCESSTYPE, nil
    }
    return 0, errors.New("Unknown SignInAccessType value: " + v)
}
func SerializeSignInAccessType(values []SignInAccessType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
