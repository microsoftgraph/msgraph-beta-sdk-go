package models
// Possible values for LocalSecurityOptionsInformationDisplayedOnLockScreen
type LocalSecurityOptionsInformationDisplayedOnLockScreenType int

const (
    // Not Configured
    NOTCONFIGURED_LOCALSECURITYOPTIONSINFORMATIONDISPLAYEDONLOCKSCREENTYPE LocalSecurityOptionsInformationDisplayedOnLockScreenType = iota
    // User display name, domain and user names
    ADMINISTRATORS_LOCALSECURITYOPTIONSINFORMATIONDISPLAYEDONLOCKSCREENTYPE
    // User display name only
    ADMINISTRATORSANDPOWERUSERS_LOCALSECURITYOPTIONSINFORMATIONDISPLAYEDONLOCKSCREENTYPE
    // Do not display user information
    ADMINISTRATORSANDINTERACTIVEUSERS_LOCALSECURITYOPTIONSINFORMATIONDISPLAYEDONLOCKSCREENTYPE
)

func (i LocalSecurityOptionsInformationDisplayedOnLockScreenType) String() string {
    return []string{"notConfigured", "administrators", "administratorsAndPowerUsers", "administratorsAndInteractiveUsers"}[i]
}
func ParseLocalSecurityOptionsInformationDisplayedOnLockScreenType(v string) (any, error) {
    result := NOTCONFIGURED_LOCALSECURITYOPTIONSINFORMATIONDISPLAYEDONLOCKSCREENTYPE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_LOCALSECURITYOPTIONSINFORMATIONDISPLAYEDONLOCKSCREENTYPE
        case "administrators":
            result = ADMINISTRATORS_LOCALSECURITYOPTIONSINFORMATIONDISPLAYEDONLOCKSCREENTYPE
        case "administratorsAndPowerUsers":
            result = ADMINISTRATORSANDPOWERUSERS_LOCALSECURITYOPTIONSINFORMATIONDISPLAYEDONLOCKSCREENTYPE
        case "administratorsAndInteractiveUsers":
            result = ADMINISTRATORSANDINTERACTIVEUSERS_LOCALSECURITYOPTIONSINFORMATIONDISPLAYEDONLOCKSCREENTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeLocalSecurityOptionsInformationDisplayedOnLockScreenType(values []LocalSecurityOptionsInformationDisplayedOnLockScreenType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i LocalSecurityOptionsInformationDisplayedOnLockScreenType) isMultiValue() bool {
    return false
}
