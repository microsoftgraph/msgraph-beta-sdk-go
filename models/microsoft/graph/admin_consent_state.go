package graph
import (
    "strings"
    "errors"
)
// 
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
    result := NOTCONFIGURED_ADMINCONSENTSTATE
    switch strings.ToUpper(v) {
        case "NOTCONFIGURED":
            result = NOTCONFIGURED_ADMINCONSENTSTATE
        case "GRANTED":
            result = GRANTED_ADMINCONSENTSTATE
        case "NOTGRANTED":
            result = NOTGRANTED_ADMINCONSENTSTATE
        default:
            return 0, errors.New("Unknown AdminConsentState value: " + v)
    }
    return &result, nil
}
func SerializeAdminConsentState(values []AdminConsentState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
