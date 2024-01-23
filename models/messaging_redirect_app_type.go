package models
import (
    "errors"
)
// Defines how app messaging redirection is protected by an App Protection Policy. Default is anyApp.
type MessagingRedirectAppType int

const (
    // App protection policy will allow messaging redirection to any app.
    ANYAPP_MESSAGINGREDIRECTAPPTYPE MessagingRedirectAppType = iota
    // App protection policy will allow messaging redirection to any managed application.
    ANYMANAGEDAPP_MESSAGINGREDIRECTAPPTYPE
    // App protection policy will allow messaging redirection only to specified applications in related App protection policy settings. See related settings `messagingRedirectAppDisplayName`, `messagingRedirectAppPackageId` and `messagingRedirectAppUrlScheme`.
    SPECIFICAPPS_MESSAGINGREDIRECTAPPTYPE
    // App protection policy will block messaging redirection to any app.
    BLOCKED_MESSAGINGREDIRECTAPPTYPE
)

func (i MessagingRedirectAppType) String() string {
    return []string{"anyApp", "anyManagedApp", "specificApps", "blocked"}[i]
}
func ParseMessagingRedirectAppType(v string) (any, error) {
    result := ANYAPP_MESSAGINGREDIRECTAPPTYPE
    switch v {
        case "anyApp":
            result = ANYAPP_MESSAGINGREDIRECTAPPTYPE
        case "anyManagedApp":
            result = ANYMANAGEDAPP_MESSAGINGREDIRECTAPPTYPE
        case "specificApps":
            result = SPECIFICAPPS_MESSAGINGREDIRECTAPPTYPE
        case "blocked":
            result = BLOCKED_MESSAGINGREDIRECTAPPTYPE
        default:
            return 0, errors.New("Unknown MessagingRedirectAppType value: " + v)
    }
    return &result, nil
}
func SerializeMessagingRedirectAppType(values []MessagingRedirectAppType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MessagingRedirectAppType) isMultiValue() bool {
    return false
}
