package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type CloudPcDeviceImageOsStatus int

const (
    SUPPORTED_CLOUDPCDEVICEIMAGEOSSTATUS CloudPcDeviceImageOsStatus = iota
    SUPPORTEDWITHWARNING_CLOUDPCDEVICEIMAGEOSSTATUS
    UNKNOWNFUTUREVALUE_CLOUDPCDEVICEIMAGEOSSTATUS
)

func (i CloudPcDeviceImageOsStatus) String() string {
    return []string{"SUPPORTED", "SUPPORTEDWITHWARNING", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseCloudPcDeviceImageOsStatus(v string) (interface{}, error) {
    result := SUPPORTED_CLOUDPCDEVICEIMAGEOSSTATUS
    switch strings.ToUpper(v) {
        case "SUPPORTED":
            result = SUPPORTED_CLOUDPCDEVICEIMAGEOSSTATUS
        case "SUPPORTEDWITHWARNING":
            result = SUPPORTEDWITHWARNING_CLOUDPCDEVICEIMAGEOSSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_CLOUDPCDEVICEIMAGEOSSTATUS
        default:
            return 0, errors.New("Unknown CloudPcDeviceImageOsStatus value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcDeviceImageOsStatus(values []CloudPcDeviceImageOsStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
