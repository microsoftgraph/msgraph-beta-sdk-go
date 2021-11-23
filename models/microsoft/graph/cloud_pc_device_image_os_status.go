package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "SUPPORTED":
            return SUPPORTED_CLOUDPCDEVICEIMAGEOSSTATUS, nil
        case "SUPPORTEDWITHWARNING":
            return SUPPORTEDWITHWARNING_CLOUDPCDEVICEIMAGEOSSTATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_CLOUDPCDEVICEIMAGEOSSTATUS, nil
    }
    return 0, errors.New("Unknown CloudPcDeviceImageOsStatus value: " + v)
}
func SerializeCloudPcDeviceImageOsStatus(values []CloudPcDeviceImageOsStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
