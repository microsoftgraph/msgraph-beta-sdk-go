package managedtenants
import (
    "errors"
    "strings"
)
// 
type NotificationDestination int

const (
    NONE_NOTIFICATIONDESTINATION NotificationDestination = iota
    API_NOTIFICATIONDESTINATION
    EMAIL_NOTIFICATIONDESTINATION
    SMS_NOTIFICATIONDESTINATION
    UNKNOWNFUTUREVALUE_NOTIFICATIONDESTINATION
)

func (i NotificationDestination) String() string {
    var values []string
    for p := NotificationDestination(1); p <= UNKNOWNFUTUREVALUE_NOTIFICATIONDESTINATION; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"none", "api", "email", "sms", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseNotificationDestination(v string) (any, error) {
    var result NotificationDestination
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_NOTIFICATIONDESTINATION
            case "api":
                result |= API_NOTIFICATIONDESTINATION
            case "email":
                result |= EMAIL_NOTIFICATIONDESTINATION
            case "sms":
                result |= SMS_NOTIFICATIONDESTINATION
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_NOTIFICATIONDESTINATION
            default:
                return 0, errors.New("Unknown NotificationDestination value: " + v)
        }
    }
    return &result, nil
}
func SerializeNotificationDestination(values []NotificationDestination) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i NotificationDestination) isMultiValue() bool {
    return true
}
