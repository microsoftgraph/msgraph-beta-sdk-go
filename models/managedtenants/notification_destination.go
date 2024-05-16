package managedtenants
import (
    "math"
    "strings"
)
type NotificationDestination int

const (
    NONE_NOTIFICATIONDESTINATION = 1
    API_NOTIFICATIONDESTINATION = 2
    EMAIL_NOTIFICATIONDESTINATION = 4
    SMS_NOTIFICATIONDESTINATION = 8
    UNKNOWNFUTUREVALUE_NOTIFICATIONDESTINATION = 16
)

func (i NotificationDestination) String() string {
    var values []string
    options := []string{"none", "api", "email", "sms", "unknownFutureValue"}
    for p := 0; p < 5; p++ {
        mantis := NotificationDestination(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
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
                return nil, nil
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
