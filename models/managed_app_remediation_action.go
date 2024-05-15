package models
// An admin initiated action to be applied on a managed app.
type ManagedAppRemediationAction int

const (
    // app and the corresponding company data to be blocked
    BLOCK_MANAGEDAPPREMEDIATIONACTION ManagedAppRemediationAction = iota
    // app and the corresponding company data to be wiped
    WIPE_MANAGEDAPPREMEDIATIONACTION
    // app and the corresponding user to be warned
    WARN_MANAGEDAPPREMEDIATIONACTION
)

func (i ManagedAppRemediationAction) String() string {
    return []string{"block", "wipe", "warn"}[i]
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
