package models
import (
    "math"
    "strings"
)
type SignInAccessType int

const (
    NONE_SIGNINACCESSTYPE = 1
    B2BCOLLABORATION_SIGNINACCESSTYPE = 2
    B2BDIRECTCONNECT_SIGNINACCESSTYPE = 4
    MICROSOFTSUPPORT_SIGNINACCESSTYPE = 8
    SERVICEPROVIDER_SIGNINACCESSTYPE = 16
    UNKNOWNFUTUREVALUE_SIGNINACCESSTYPE = 32
    PASSTHROUGH_SIGNINACCESSTYPE = 64
)

func (i SignInAccessType) String() string {
    var values []string
    options := []string{"none", "b2bCollaboration", "b2bDirectConnect", "microsoftSupport", "serviceProvider", "unknownFutureValue", "passthrough"}
    for p := 0; p < 7; p++ {
        mantis := SignInAccessType(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
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
            case "passthrough":
                result |= PASSTHROUGH_SIGNINACCESSTYPE
            default:
                return nil, nil
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
