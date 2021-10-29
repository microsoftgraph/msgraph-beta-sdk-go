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
    switch strings.ToUpper(v) {
        case "NOTEVALUATED":
            return NOTEVALUATED_EMBEDDEDSIMDEVICESTATEVALUE, nil
        case "FAILED":
            return FAILED_EMBEDDEDSIMDEVICESTATEVALUE, nil
        case "INSTALLING":
            return INSTALLING_EMBEDDEDSIMDEVICESTATEVALUE, nil
        case "INSTALLED":
            return INSTALLED_EMBEDDEDSIMDEVICESTATEVALUE, nil
        case "DELETING":
            return DELETING_EMBEDDEDSIMDEVICESTATEVALUE, nil
        case "ERROR":
            return ERROR_EMBEDDEDSIMDEVICESTATEVALUE, nil
        case "DELETED":
            return DELETED_EMBEDDEDSIMDEVICESTATEVALUE, nil
        case "REMOVEDBYUSER":
            return REMOVEDBYUSER_EMBEDDEDSIMDEVICESTATEVALUE, nil
    }
    return 0, errors.New("Unknown EmbeddedSIMDeviceStateValue value: " + v)
}
func SerializeEmbeddedSIMDeviceStateValue(values []EmbeddedSIMDeviceStateValue) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
