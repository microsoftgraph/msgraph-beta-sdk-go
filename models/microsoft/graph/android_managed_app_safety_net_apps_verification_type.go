package graph
import (
    "strings"
    "errors"
)
type AndroidManagedAppSafetyNetAppsVerificationType int

const (
    NONE_ANDROIDMANAGEDAPPSAFETYNETAPPSVERIFICATIONTYPE AndroidManagedAppSafetyNetAppsVerificationType = iota
    ENABLED_ANDROIDMANAGEDAPPSAFETYNETAPPSVERIFICATIONTYPE
)

func (i AndroidManagedAppSafetyNetAppsVerificationType) String() string {
    return []string{"NONE", "ENABLED"}[i]
}
func ParseAndroidManagedAppSafetyNetAppsVerificationType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_ANDROIDMANAGEDAPPSAFETYNETAPPSVERIFICATIONTYPE, nil
        case "ENABLED":
            return ENABLED_ANDROIDMANAGEDAPPSAFETYNETAPPSVERIFICATIONTYPE, nil
    }
    return 0, errors.New("Unknown AndroidManagedAppSafetyNetAppsVerificationType value: " + v)
}
func SerializeAndroidManagedAppSafetyNetAppsVerificationType(values []AndroidManagedAppSafetyNetAppsVerificationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
