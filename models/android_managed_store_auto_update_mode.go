package models
import (
    "errors"
)
// Prioritization for automatic updates of Android Managed Store apps set on assignment.
type AndroidManagedStoreAutoUpdateMode int

const (
    // Default update behavior (device must be connected to Wifi, charging and not actively used).
    DEFAULTESCAPED_ANDROIDMANAGEDSTOREAUTOUPDATEMODE AndroidManagedStoreAutoUpdateMode = iota
    // Updates are postponed for a maximum of 90 days after the app becomes out of date.
    POSTPONED_ANDROIDMANAGEDSTOREAUTOUPDATEMODE
    // The app is updated as soon as possible by the developer. If device is online, it will be updated within minutes.
    PRIORITY_ANDROIDMANAGEDSTOREAUTOUPDATEMODE
    // Unknown future mode (reserved, not used right now).
    UNKNOWNFUTUREVALUE_ANDROIDMANAGEDSTOREAUTOUPDATEMODE
)

func (i AndroidManagedStoreAutoUpdateMode) String() string {
    return []string{"default", "postponed", "priority", "unknownFutureValue"}[i]
}
func ParseAndroidManagedStoreAutoUpdateMode(v string) (any, error) {
    result := DEFAULTESCAPED_ANDROIDMANAGEDSTOREAUTOUPDATEMODE
    switch v {
        case "default":
            result = DEFAULTESCAPED_ANDROIDMANAGEDSTOREAUTOUPDATEMODE
        case "postponed":
            result = POSTPONED_ANDROIDMANAGEDSTOREAUTOUPDATEMODE
        case "priority":
            result = PRIORITY_ANDROIDMANAGEDSTOREAUTOUPDATEMODE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ANDROIDMANAGEDSTOREAUTOUPDATEMODE
        default:
            return 0, errors.New("Unknown AndroidManagedStoreAutoUpdateMode value: " + v)
    }
    return &result, nil
}
func SerializeAndroidManagedStoreAutoUpdateMode(values []AndroidManagedStoreAutoUpdateMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
