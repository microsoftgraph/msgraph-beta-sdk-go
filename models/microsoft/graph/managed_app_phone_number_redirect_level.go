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
    switch strings.ToUpper(v) {
        case "ALLAPPS":
            return ALLAPPS_MANAGEDAPPPHONENUMBERREDIRECTLEVEL, nil
        case "MANAGEDAPPS":
            return MANAGEDAPPS_MANAGEDAPPPHONENUMBERREDIRECTLEVEL, nil
        case "CUSTOMAPP":
            return CUSTOMAPP_MANAGEDAPPPHONENUMBERREDIRECTLEVEL, nil
        case "BLOCKED":
            return BLOCKED_MANAGEDAPPPHONENUMBERREDIRECTLEVEL, nil
    }
    return 0, errors.New("Unknown ManagedAppPhoneNumberRedirectLevel value: " + v)
}
func SerializeManagedAppPhoneNumberRedirectLevel(values []ManagedAppPhoneNumberRedirectLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
