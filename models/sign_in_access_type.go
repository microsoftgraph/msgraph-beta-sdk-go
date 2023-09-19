package models
import (
    "errors"
    "strings"
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
    var values []string
    for p := SignInAccessType(1); p <= UNKNOWNFUTUREVALUE_SIGNINACCESSTYPE; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"none", "b2bCollaboration", "b2bDirectConnect", "microsoftSupport", "serviceProvider", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseSignInAccessType(v string) (any, error) {
    var result SignInAccessType
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_SIGNINACCESSTYPE
            case "b2bCollaboration":
                result |= B2BCOLLABORATION_SIGNINACCESSTYPE
            case "b2bDirectConnect":
                result |= B2BDIRECTCONNECT_SIGNINACCESSTYPE
            case "microsoftSupport":
                result |= MICROSOFTSUPPORT_SIGNINACCESSTYPE
            case "serviceProvider":
                result |= SERVICEPROVIDER_SIGNINACCESSTYPE
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_SIGNINACCESSTYPE
            default:
                return 0, errors.New("Unknown SignInAccessType value: " + v)
        }
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
func (i SignInAccessType) isMultiValue() bool {
    return true
}
