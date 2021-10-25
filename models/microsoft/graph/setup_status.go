package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_SETUPSTATUS, nil
        case "NOTREGISTEREDYET":
            return NOTREGISTEREDYET_SETUPSTATUS, nil
        case "REGISTEREDSETUPNOTSTARTED":
            return REGISTEREDSETUPNOTSTARTED_SETUPSTATUS, nil
        case "REGISTEREDSETUPINPROGRESS":
            return REGISTEREDSETUPINPROGRESS_SETUPSTATUS, nil
        case "REGISTRATIONANDSETUPCOMPLETED":
            return REGISTRATIONANDSETUPCOMPLETED_SETUPSTATUS, nil
        case "REGISTRATIONFAILED":
            return REGISTRATIONFAILED_SETUPSTATUS, nil
        case "REGISTRATIONTIMEDOUT":
            return REGISTRATIONTIMEDOUT_SETUPSTATUS, nil
        case "DISABLED":
            return DISABLED_SETUPSTATUS, nil
    }
    return 0, errors.New("Unknown SetupStatus value: " + v)
}
func SerializeSetupStatus(values []SetupStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
