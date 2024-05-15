package models
// Admin consent state.
type AdminConsentState int

const (
    // Admin did not configure the item
    NOTCONFIGURED_ADMINCONSENTSTATE AdminConsentState = iota
    // Admin granted item
    GRANTED_ADMINCONSENTSTATE
    // Admin deos not grant item
    NOTGRANTED_ADMINCONSENTSTATE
)

func (i AdminConsentState) String() string {
    return []string{"notConfigured", "granted", "notGranted"}[i]
}
func ParseAdminConsentState(v string) (any, error) {
    result := NOTCONFIGURED_ADMINCONSENTSTATE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_ADMINCONSENTSTATE
        case "granted":
            result = GRANTED_ADMINCONSENTSTATE
        case "notGranted":
            result = NOTGRANTED_ADMINCONSENTSTATE
        default:
            return nil, nil
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
func (i AdminConsentState) isMultiValue() bool {
    return false
}
