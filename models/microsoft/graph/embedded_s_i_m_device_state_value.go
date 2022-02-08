package graph
import (
    "strings"
    "errors"
)
// 
type EmbeddedSIMDeviceStateValue int

const (
    NOTEVALUATED_EMBEDDEDSIMDEVICESTATEVALUE EmbeddedSIMDeviceStateValue = iota
    FAILED_EMBEDDEDSIMDEVICESTATEVALUE
    INSTALLING_EMBEDDEDSIMDEVICESTATEVALUE
    INSTALLED_EMBEDDEDSIMDEVICESTATEVALUE
    DELETING_EMBEDDEDSIMDEVICESTATEVALUE
    ERROR_EMBEDDEDSIMDEVICESTATEVALUE
    DELETED_EMBEDDEDSIMDEVICESTATEVALUE
    REMOVEDBYUSER_EMBEDDEDSIMDEVICESTATEVALUE
)

func (i EmbeddedSIMDeviceStateValue) String() string {
    return []string{"NOTEVALUATED", "FAILED", "INSTALLING", "INSTALLED", "DELETING", "ERROR", "DELETED", "REMOVEDBYUSER"}[i]
}
func ParseEmbeddedSIMDeviceStateValue(v string) (interface{}, error) {
    result := NOTEVALUATED_EMBEDDEDSIMDEVICESTATEVALUE
    switch strings.ToUpper(v) {
        case "NOTEVALUATED":
            result = NOTEVALUATED_EMBEDDEDSIMDEVICESTATEVALUE
        case "FAILED":
            result = FAILED_EMBEDDEDSIMDEVICESTATEVALUE
        case "INSTALLING":
            result = INSTALLING_EMBEDDEDSIMDEVICESTATEVALUE
        case "INSTALLED":
            result = INSTALLED_EMBEDDEDSIMDEVICESTATEVALUE
        case "DELETING":
            result = DELETING_EMBEDDEDSIMDEVICESTATEVALUE
        case "ERROR":
            result = ERROR_EMBEDDEDSIMDEVICESTATEVALUE
        case "DELETED":
            result = DELETED_EMBEDDEDSIMDEVICESTATEVALUE
        case "REMOVEDBYUSER":
            result = REMOVEDBYUSER_EMBEDDEDSIMDEVICESTATEVALUE
        default:
            return 0, errors.New("Unknown EmbeddedSIMDeviceStateValue value: " + v)
    }
    return &result, nil
}
func SerializeEmbeddedSIMDeviceStateValue(values []EmbeddedSIMDeviceStateValue) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
