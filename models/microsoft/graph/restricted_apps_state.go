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
    switch strings.ToUpper(v) {
        case "PROHIBITEDAPPS":
            return PROHIBITEDAPPS_RESTRICTEDAPPSSTATE, nil
        case "NOTAPPROVEDAPPS":
            return NOTAPPROVEDAPPS_RESTRICTEDAPPSSTATE, nil
    }
    return 0, errors.New("Unknown RestrictedAppsState value: " + v)
}
func SerializeRestrictedAppsState(values []RestrictedAppsState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
