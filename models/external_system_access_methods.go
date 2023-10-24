package models
import (
    "errors"
    "strings"
)
// 
type ExternalSystemAccessMethods int

const (
    DIRECT_EXTERNALSYSTEMACCESSMETHODS ExternalSystemAccessMethods = iota
    ROLECHAINING_EXTERNALSYSTEMACCESSMETHODS
    UNKNOWNFUTUREVALUE_EXTERNALSYSTEMACCESSMETHODS
)

func (i ExternalSystemAccessMethods) String() string {
    var values []string
    for p := ExternalSystemAccessMethods(1); p <= UNKNOWNFUTUREVALUE_EXTERNALSYSTEMACCESSMETHODS; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"direct", "roleChaining", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseExternalSystemAccessMethods(v string) (any, error) {
    var result ExternalSystemAccessMethods
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "direct":
                result |= DIRECT_EXTERNALSYSTEMACCESSMETHODS
            case "roleChaining":
                result |= ROLECHAINING_EXTERNALSYSTEMACCESSMETHODS
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_EXTERNALSYSTEMACCESSMETHODS
            default:
                return 0, errors.New("Unknown ExternalSystemAccessMethods value: " + v)
        }
    }
    return &result, nil
}
func SerializeExternalSystemAccessMethods(values []ExternalSystemAccessMethods) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ExternalSystemAccessMethods) isMultiValue() bool {
    return true
}
