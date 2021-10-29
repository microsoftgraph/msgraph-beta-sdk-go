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
    switch strings.ToUpper(v) {
        case "REGISTERED":
            return REGISTERED_REGISTRATIONSTATUSTYPE, nil
        case "ENABLED":
            return ENABLED_REGISTRATIONSTATUSTYPE, nil
        case "CAPABLE":
            return CAPABLE_REGISTRATIONSTATUSTYPE, nil
        case "MFAREGISTERED":
            return MFAREGISTERED_REGISTRATIONSTATUSTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_REGISTRATIONSTATUSTYPE, nil
    }
    return 0, errors.New("Unknown RegistrationStatusType value: " + v)
}
func SerializeRegistrationStatusType(values []RegistrationStatusType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
