package graph
import (
    "strings"
    "errors"
)
// 
type ManagedAppPhoneNumberRedirectLevel int

const (
    ALLAPPS_MANAGEDAPPPHONENUMBERREDIRECTLEVEL ManagedAppPhoneNumberRedirectLevel = iota
    MANAGEDAPPS_MANAGEDAPPPHONENUMBERREDIRECTLEVEL
    CUSTOMAPP_MANAGEDAPPPHONENUMBERREDIRECTLEVEL
    BLOCKED_MANAGEDAPPPHONENUMBERREDIRECTLEVEL
)

func (i ManagedAppPhoneNumberRedirectLevel) String() string {
    return []string{"ALLAPPS", "MANAGEDAPPS", "CUSTOMAPP", "BLOCKED"}[i]
}
func ParseManagedAppPhoneNumberRedirectLevel(v string) (interface{}, error) {
    result := ALLAPPS_MANAGEDAPPPHONENUMBERREDIRECTLEVEL
    switch strings.ToUpper(v) {
        case "ALLAPPS":
            result = ALLAPPS_MANAGEDAPPPHONENUMBERREDIRECTLEVEL
        case "MANAGEDAPPS":
            result = MANAGEDAPPS_MANAGEDAPPPHONENUMBERREDIRECTLEVEL
        case "CUSTOMAPP":
            result = CUSTOMAPP_MANAGEDAPPPHONENUMBERREDIRECTLEVEL
        case "BLOCKED":
            result = BLOCKED_MANAGEDAPPPHONENUMBERREDIRECTLEVEL
        default:
            return 0, errors.New("Unknown ManagedAppPhoneNumberRedirectLevel value: " + v)
    }
    return &result, nil
}
func SerializeManagedAppPhoneNumberRedirectLevel(values []ManagedAppPhoneNumberRedirectLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
