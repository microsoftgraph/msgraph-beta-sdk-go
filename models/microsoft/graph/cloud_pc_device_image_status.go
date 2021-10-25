package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "PENDING":
            return PENDING_CLOUDPCDEVICEIMAGESTATUS, nil
        case "READY":
            return READY_CLOUDPCDEVICEIMAGESTATUS, nil
        case "FAILED":
            return FAILED_CLOUDPCDEVICEIMAGESTATUS, nil
    }
    return 0, errors.New("Unknown CloudPcDeviceImageStatus value: " + v)
}
func SerializeCloudPcDeviceImageStatus(values []CloudPcDeviceImageStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
