package graph
import (
    "strings"
    "errors"
)
type AdminConsentState int

const (
    NOTCONFIGURED_ADMINCONSENTSTATE AdminConsentState = iota
    GRANTED_ADMINCONSENTSTATE
    NOTGRANTED_ADMINCONSENTSTATE
)

func (i AdminConsentState) String() string {
    return []string{"NOTCONFIGURED", "GRANTED", "NOTGRANTED"}[i]
}
func ParseAdminConsentState(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NOTCONFIGURED":
            return NOTCONFIGURED_ADMINCONSENTSTATE, nil
        case "GRANTED":
            return GRANTED_ADMINCONSENTSTATE, nil
        case "NOTGRANTED":
            return NOTGRANTED_ADMINCONSENTSTATE, nil
    }
    return 0, errors.New("Unknown AdminConsentState value: " + v)
}
func SerializeAdminConsentState(values []AdminConsentState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
