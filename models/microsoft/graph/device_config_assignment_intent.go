package graph
import (
    "strings"
    "errors"
)
// 
type DeviceConfigAssignmentIntent int

const (
    APPLY_DEVICECONFIGASSIGNMENTINTENT DeviceConfigAssignmentIntent = iota
    REMOVE_DEVICECONFIGASSIGNMENTINTENT
)

func (i DeviceConfigAssignmentIntent) String() string {
    return []string{"APPLY", "REMOVE"}[i]
}
func ParseDeviceConfigAssignmentIntent(v string) (interface{}, error) {
    result := APPLY_DEVICECONFIGASSIGNMENTINTENT
    switch strings.ToUpper(v) {
        case "APPLY":
            result = APPLY_DEVICECONFIGASSIGNMENTINTENT
        case "REMOVE":
            result = REMOVE_DEVICECONFIGASSIGNMENTINTENT
        default:
            return 0, errors.New("Unknown DeviceConfigAssignmentIntent value: " + v)
    }
    return &result, nil
}
func SerializeDeviceConfigAssignmentIntent(values []DeviceConfigAssignmentIntent) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
