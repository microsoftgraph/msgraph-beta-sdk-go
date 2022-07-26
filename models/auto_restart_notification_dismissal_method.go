package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type AutoRestartNotificationDismissalMethod int

const (
    // Not configured
    NOTCONFIGURED_AUTORESTARTNOTIFICATIONDISMISSALMETHOD AutoRestartNotificationDismissalMethod = iota
    // Auto dismissal
    AUTOMATIC_AUTORESTARTNOTIFICATIONDISMISSALMETHOD
    // User dismissal
    USER_AUTORESTARTNOTIFICATIONDISMISSALMETHOD
)

func (i AutoRestartNotificationDismissalMethod) String() string {
    return []string{"notConfigured", "automatic", "user"}[i]
}
func ParseAutoRestartNotificationDismissalMethod(v string) (interface{}, error) {
    result := NOTCONFIGURED_AUTORESTARTNOTIFICATIONDISMISSALMETHOD
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_AUTORESTARTNOTIFICATIONDISMISSALMETHOD
        case "automatic":
            result = AUTOMATIC_AUTORESTARTNOTIFICATIONDISMISSALMETHOD
        case "user":
            result = USER_AUTORESTARTNOTIFICATIONDISMISSALMETHOD
        default:
            return 0, errors.New("Unknown AutoRestartNotificationDismissalMethod value: " + v)
    }
    return &result, nil
}
func SerializeAutoRestartNotificationDismissalMethod(values []AutoRestartNotificationDismissalMethod) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
