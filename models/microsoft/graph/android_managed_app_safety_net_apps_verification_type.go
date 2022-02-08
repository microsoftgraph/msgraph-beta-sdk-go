package graph
import (
    "strings"
    "errors"
)
// 
type AndroidManagedAppSafetyNetAppsVerificationType int

const (
    NONE_ANDROIDMANAGEDAPPSAFETYNETAPPSVERIFICATIONTYPE AndroidManagedAppSafetyNetAppsVerificationType = iota
    ENABLED_ANDROIDMANAGEDAPPSAFETYNETAPPSVERIFICATIONTYPE
)

func (i AndroidManagedAppSafetyNetAppsVerificationType) String() string {
    return []string{"NONE", "ENABLED"}[i]
}
func ParseAndroidManagedAppSafetyNetAppsVerificationType(v string) (interface{}, error) {
    result := NONE_ANDROIDMANAGEDAPPSAFETYNETAPPSVERIFICATIONTYPE
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_ANDROIDMANAGEDAPPSAFETYNETAPPSVERIFICATIONTYPE
        case "ENABLED":
            result = ENABLED_ANDROIDMANAGEDAPPSAFETYNETAPPSVERIFICATIONTYPE
        default:
            return 0, errors.New("Unknown AndroidManagedAppSafetyNetAppsVerificationType value: " + v)
    }
    return &result, nil
}
func SerializeAndroidManagedAppSafetyNetAppsVerificationType(values []AndroidManagedAppSafetyNetAppsVerificationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
