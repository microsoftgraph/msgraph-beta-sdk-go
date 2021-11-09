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
    switch strings.ToUpper(v) {
        case "APPLY":
            return APPLY_DEVICECONFIGASSIGNMENTINTENT, nil
        case "REMOVE":
            return REMOVE_DEVICECONFIGASSIGNMENTINTENT, nil
    }
    return 0, errors.New("Unknown DeviceConfigAssignmentIntent value: " + v)
}
func SerializeDeviceConfigAssignmentIntent(values []DeviceConfigAssignmentIntent) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
