package models
// Restricted apps state
type RestrictedAppsState int

const (
    // Prohibited apps
    PROHIBITEDAPPS_RESTRICTEDAPPSSTATE RestrictedAppsState = iota
    // Not approved apps
    NOTAPPROVEDAPPS_RESTRICTEDAPPSSTATE
)

func (i RestrictedAppsState) String() string {
    return []string{"prohibitedApps", "notApprovedApps"}[i]
}
func ParseRestrictedAppsState(v string) (any, error) {
    result := PROHIBITEDAPPS_RESTRICTEDAPPSSTATE
    switch v {
        case "prohibitedApps":
            result = PROHIBITEDAPPS_RESTRICTEDAPPSSTATE
        case "notApprovedApps":
            result = NOTAPPROVEDAPPS_RESTRICTEDAPPSSTATE
        default:
            return nil, nil
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
func (i RestrictedAppsState) isMultiValue() bool {
    return false
}
