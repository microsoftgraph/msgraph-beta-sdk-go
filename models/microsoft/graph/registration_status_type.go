package graph
import (
    "strings"
    "errors"
)
// 
type RegistrationStatusType int

const (
    REGISTERED_REGISTRATIONSTATUSTYPE RegistrationStatusType = iota
    ENABLED_REGISTRATIONSTATUSTYPE
    CAPABLE_REGISTRATIONSTATUSTYPE
    MFAREGISTERED_REGISTRATIONSTATUSTYPE
    UNKNOWNFUTUREVALUE_REGISTRATIONSTATUSTYPE
)

func (i RegistrationStatusType) String() string {
    return []string{"REGISTERED", "ENABLED", "CAPABLE", "MFAREGISTERED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseRegistrationStatusType(v string) (interface{}, error) {
    result := REGISTERED_REGISTRATIONSTATUSTYPE
    switch strings.ToUpper(v) {
        case "REGISTERED":
            result = REGISTERED_REGISTRATIONSTATUSTYPE
        case "ENABLED":
            result = ENABLED_REGISTRATIONSTATUSTYPE
        case "CAPABLE":
            result = CAPABLE_REGISTRATIONSTATUSTYPE
        case "MFAREGISTERED":
            result = MFAREGISTERED_REGISTRATIONSTATUSTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_REGISTRATIONSTATUSTYPE
        default:
            return 0, errors.New("Unknown RegistrationStatusType value: " + v)
    }
    return &result, nil
}
func SerializeRegistrationStatusType(values []RegistrationStatusType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
