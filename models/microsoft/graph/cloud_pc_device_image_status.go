package graph
import (
    "strings"
    "errors"
)
// 
type CloudPcDeviceImageStatus int

const (
    PENDING_CLOUDPCDEVICEIMAGESTATUS CloudPcDeviceImageStatus = iota
    READY_CLOUDPCDEVICEIMAGESTATUS
    FAILED_CLOUDPCDEVICEIMAGESTATUS
)

func (i CloudPcDeviceImageStatus) String() string {
    return []string{"PENDING", "READY", "FAILED"}[i]
}
func ParseCloudPcDeviceImageStatus(v string) (interface{}, error) {
    result := PENDING_CLOUDPCDEVICEIMAGESTATUS
    switch strings.ToUpper(v) {
        case "PENDING":
            result = PENDING_CLOUDPCDEVICEIMAGESTATUS
        case "READY":
            result = READY_CLOUDPCDEVICEIMAGESTATUS
        case "FAILED":
            result = FAILED_CLOUDPCDEVICEIMAGESTATUS
        default:
            return 0, errors.New("Unknown CloudPcDeviceImageStatus value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcDeviceImageStatus(values []CloudPcDeviceImageStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
