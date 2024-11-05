package models
// An admin initiated action to be applied on a managed app.
type ManagedAppRemediationAction int

const (
    // Indicates the user will be blocked from accessing the app and corporate data
    BLOCK_MANAGEDAPPREMEDIATIONACTION ManagedAppRemediationAction = iota
    // Indicates the corporate data will be removed from the app
    WIPE_MANAGEDAPPREMEDIATIONACTION
    // Indicates user will be warned the when accessing the app
    WARN_MANAGEDAPPREMEDIATIONACTION
    // Indicates user will be blocked from accessing the app and corporate data if devices supports this setting
    BLOCKWHENSETTINGISSUPPORTED_MANAGEDAPPREMEDIATIONACTION
)

func (i ManagedAppRemediationAction) String() string {
    return []string{"block", "wipe", "warn", "blockWhenSettingIsSupported"}[i]
}
func ParseManagedAppRemediationAction(v string) (any, error) {
    result := BLOCK_MANAGEDAPPREMEDIATIONACTION
    switch v {
        case "block":
            result = BLOCK_MANAGEDAPPREMEDIATIONACTION
        case "wipe":
            result = WIPE_MANAGEDAPPREMEDIATIONACTION
        case "warn":
            result = WARN_MANAGEDAPPREMEDIATIONACTION
        case "blockWhenSettingIsSupported":
            result = BLOCKWHENSETTINGISSUPPORTED_MANAGEDAPPREMEDIATIONACTION
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeManagedAppRemediationAction(values []ManagedAppRemediationAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ManagedAppRemediationAction) isMultiValue() bool {
    return false
}
