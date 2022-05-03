package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the policyRoot singleton.
type SignInFrequencyInterval int

const (
    TIMEBASED_SIGNINFREQUENCYINTERVAL SignInFrequencyInterval = iota
    EVERYTIME_SIGNINFREQUENCYINTERVAL
    UNKNOWNFUTUREVALUE_SIGNINFREQUENCYINTERVAL
)

func (i SignInFrequencyInterval) String() string {
    return []string{"TIMEBASED", "EVERYTIME", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseSignInFrequencyInterval(v string) (interface{}, error) {
    result := TIMEBASED_SIGNINFREQUENCYINTERVAL
    switch strings.ToUpper(v) {
        case "TIMEBASED":
            result = TIMEBASED_SIGNINFREQUENCYINTERVAL
        case "EVERYTIME":
            result = EVERYTIME_SIGNINFREQUENCYINTERVAL
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_SIGNINFREQUENCYINTERVAL
        default:
            return 0, errors.New("Unknown SignInFrequencyInterval value: " + v)
    }
    return &result, nil
}
func SerializeSignInFrequencyInterval(values []SignInFrequencyInterval) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
