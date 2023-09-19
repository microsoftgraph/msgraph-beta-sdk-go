package networkaccess
import (
    "errors"
)
// 
type DeviceCategory int

const (
    CLIENT_DEVICECATEGORY DeviceCategory = iota
    BRANCH_DEVICECATEGORY
    UNKNOWNFUTUREVALUE_DEVICECATEGORY
)

func (i DeviceCategory) String() string {
    return []string{"client", "branch", "unknownFutureValue"}[i]
}
func ParseDeviceCategory(v string) (any, error) {
    result := CLIENT_DEVICECATEGORY
    switch v {
        case "client":
            result = CLIENT_DEVICECATEGORY
        case "branch":
            result = BRANCH_DEVICECATEGORY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DEVICECATEGORY
        default:
            return 0, errors.New("Unknown DeviceCategory value: " + v)
    }
    return &result, nil
}
func SerializeDeviceCategory(values []DeviceCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DeviceCategory) isMultiValue() bool {
    return false
}
