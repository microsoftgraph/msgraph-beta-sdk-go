package models
// The classes of apps that are allowed to click-to-open a phone number, for making phone calls or sending text messages.
type ManagedAppPhoneNumberRedirectLevel int

const (
    // Sharing is allowed to all apps.
    ALLAPPS_MANAGEDAPPPHONENUMBERREDIRECTLEVEL ManagedAppPhoneNumberRedirectLevel = iota
    // Sharing is allowed to all managed apps.
    MANAGEDAPPS_MANAGEDAPPPHONENUMBERREDIRECTLEVEL
    // Sharing is allowed to a custom app.
    CUSTOMAPP_MANAGEDAPPPHONENUMBERREDIRECTLEVEL
    // Sharing between apps is blocked.
    BLOCKED_MANAGEDAPPPHONENUMBERREDIRECTLEVEL
)

func (i ManagedAppPhoneNumberRedirectLevel) String() string {
    return []string{"allApps", "managedApps", "customApp", "blocked"}[i]
}
func ParseManagedAppPhoneNumberRedirectLevel(v string) (any, error) {
    result := ALLAPPS_MANAGEDAPPPHONENUMBERREDIRECTLEVEL
    switch v {
        case "allApps":
            result = ALLAPPS_MANAGEDAPPPHONENUMBERREDIRECTLEVEL
        case "managedApps":
            result = MANAGEDAPPS_MANAGEDAPPPHONENUMBERREDIRECTLEVEL
        case "customApp":
            result = CUSTOMAPP_MANAGEDAPPPHONENUMBERREDIRECTLEVEL
        case "blocked":
            result = BLOCKED_MANAGEDAPPPHONENUMBERREDIRECTLEVEL
        default:
            return nil, nil
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
func (i ManagedAppPhoneNumberRedirectLevel) isMultiValue() bool {
    return false
}
