package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of riskDetection entities.
type ActivityType int

const (
    SIGNIN_ACTIVITYTYPE ActivityType = iota
    USER_ACTIVITYTYPE
    UNKNOWNFUTUREVALUE_ACTIVITYTYPE
    SERVICEPRINCIPAL_ACTIVITYTYPE
)

func (i ActivityType) String() string {
    return []string{"SIGNIN", "USER", "UNKNOWNFUTUREVALUE", "SERVICEPRINCIPAL"}[i]
}
func ParseActivityType(v string) (interface{}, error) {
    result := SIGNIN_ACTIVITYTYPE
    switch strings.ToUpper(v) {
        case "SIGNIN":
            result = SIGNIN_ACTIVITYTYPE
        case "USER":
            result = USER_ACTIVITYTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ACTIVITYTYPE
        case "SERVICEPRINCIPAL":
            result = SERVICEPRINCIPAL_ACTIVITYTYPE
        default:
            return 0, errors.New("Unknown ActivityType value: " + v)
    }
    return &result, nil
}
func SerializeActivityType(values []ActivityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
