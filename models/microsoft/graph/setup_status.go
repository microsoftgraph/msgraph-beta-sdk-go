package graph
import (
    "strings"
    "errors"
)
// 
type SetupStatus int

const (
    UNKNOWN_SETUPSTATUS SetupStatus = iota
    NOTREGISTEREDYET_SETUPSTATUS
    REGISTEREDSETUPNOTSTARTED_SETUPSTATUS
    REGISTEREDSETUPINPROGRESS_SETUPSTATUS
    REGISTRATIONANDSETUPCOMPLETED_SETUPSTATUS
    REGISTRATIONFAILED_SETUPSTATUS
    REGISTRATIONTIMEDOUT_SETUPSTATUS
    DISABLED_SETUPSTATUS
)

func (i SetupStatus) String() string {
    return []string{"UNKNOWN", "NOTREGISTEREDYET", "REGISTEREDSETUPNOTSTARTED", "REGISTEREDSETUPINPROGRESS", "REGISTRATIONANDSETUPCOMPLETED", "REGISTRATIONFAILED", "REGISTRATIONTIMEDOUT", "DISABLED"}[i]
}
func ParseSetupStatus(v string) (interface{}, error) {
    result := UNKNOWN_SETUPSTATUS
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_SETUPSTATUS
        case "NOTREGISTEREDYET":
            result = NOTREGISTEREDYET_SETUPSTATUS
        case "REGISTEREDSETUPNOTSTARTED":
            result = REGISTEREDSETUPNOTSTARTED_SETUPSTATUS
        case "REGISTEREDSETUPINPROGRESS":
            result = REGISTEREDSETUPINPROGRESS_SETUPSTATUS
        case "REGISTRATIONANDSETUPCOMPLETED":
            result = REGISTRATIONANDSETUPCOMPLETED_SETUPSTATUS
        case "REGISTRATIONFAILED":
            result = REGISTRATIONFAILED_SETUPSTATUS
        case "REGISTRATIONTIMEDOUT":
            result = REGISTRATIONTIMEDOUT_SETUPSTATUS
        case "DISABLED":
            result = DISABLED_SETUPSTATUS
        default:
            return 0, errors.New("Unknown SetupStatus value: " + v)
    }
    return &result, nil
}
func SerializeSetupStatus(values []SetupStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
