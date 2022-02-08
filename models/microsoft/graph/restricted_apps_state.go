package graph
import (
    "strings"
    "errors"
)
// 
type RestrictedAppsState int

const (
    PROHIBITEDAPPS_RESTRICTEDAPPSSTATE RestrictedAppsState = iota
    NOTAPPROVEDAPPS_RESTRICTEDAPPSSTATE
)

func (i RestrictedAppsState) String() string {
    return []string{"PROHIBITEDAPPS", "NOTAPPROVEDAPPS"}[i]
}
func ParseRestrictedAppsState(v string) (interface{}, error) {
    result := PROHIBITEDAPPS_RESTRICTEDAPPSSTATE
    switch strings.ToUpper(v) {
        case "PROHIBITEDAPPS":
            result = PROHIBITEDAPPS_RESTRICTEDAPPSSTATE
        case "NOTAPPROVEDAPPS":
            result = NOTAPPROVEDAPPS_RESTRICTEDAPPSSTATE
        default:
            return 0, errors.New("Unknown RestrictedAppsState value: " + v)
    }
    return &result, nil
}
func SerializeRestrictedAppsState(values []RestrictedAppsState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
