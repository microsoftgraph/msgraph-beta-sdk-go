// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models
// Possible values for LocalSecurityOptionsInformationShownOnLockScreenType
type LocalSecurityOptionsInformationShownOnLockScreenType int

const (
    // Not Configured
    NOTCONFIGURED_LOCALSECURITYOPTIONSINFORMATIONSHOWNONLOCKSCREENTYPE LocalSecurityOptionsInformationShownOnLockScreenType = iota
    // User display name, domain and user names
    USERDISPLAYNAMEDOMAINUSER_LOCALSECURITYOPTIONSINFORMATIONSHOWNONLOCKSCREENTYPE
    // User display name only
    USERDISPLAYNAMEONLY_LOCALSECURITYOPTIONSINFORMATIONSHOWNONLOCKSCREENTYPE
    // Do not display user information
    DONOTDISPLAYUSER_LOCALSECURITYOPTIONSINFORMATIONSHOWNONLOCKSCREENTYPE
)

func (i LocalSecurityOptionsInformationShownOnLockScreenType) String() string {
    return []string{"notConfigured", "userDisplayNameDomainUser", "userDisplayNameOnly", "doNotDisplayUser"}[i]
}
func ParseLocalSecurityOptionsInformationShownOnLockScreenType(v string) (any, error) {
    result := NOTCONFIGURED_LOCALSECURITYOPTIONSINFORMATIONSHOWNONLOCKSCREENTYPE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_LOCALSECURITYOPTIONSINFORMATIONSHOWNONLOCKSCREENTYPE
        case "userDisplayNameDomainUser":
            result = USERDISPLAYNAMEDOMAINUSER_LOCALSECURITYOPTIONSINFORMATIONSHOWNONLOCKSCREENTYPE
        case "userDisplayNameOnly":
            result = USERDISPLAYNAMEONLY_LOCALSECURITYOPTIONSINFORMATIONSHOWNONLOCKSCREENTYPE
        case "doNotDisplayUser":
            result = DONOTDISPLAYUSER_LOCALSECURITYOPTIONSINFORMATIONSHOWNONLOCKSCREENTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeLocalSecurityOptionsInformationShownOnLockScreenType(values []LocalSecurityOptionsInformationShownOnLockScreenType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i LocalSecurityOptionsInformationShownOnLockScreenType) isMultiValue() bool {
    return false
}
