package security
import (
    "errors"
)
// 
type DeviceIdEntityIdentifier int

const (
    DEVICEID_DEVICEIDENTITYIDENTIFIER DeviceIdEntityIdentifier = iota
    UNKNOWNFUTUREVALUE_DEVICEIDENTITYIDENTIFIER
)

func (i DeviceIdEntityIdentifier) String() string {
    return []string{"deviceId", "unknownFutureValue"}[i]
}
func ParseDeviceIdEntityIdentifier(v string) (any, error) {
    result := DEVICEID_DEVICEIDENTITYIDENTIFIER
    switch v {
        case "deviceId":
            result = DEVICEID_DEVICEIDENTITYIDENTIFIER
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DEVICEIDENTITYIDENTIFIER
        default:
            return 0, errors.New("Unknown DeviceIdEntityIdentifier value: " + v)
    }
    return &result, nil
}
func SerializeDeviceIdEntityIdentifier(values []DeviceIdEntityIdentifier) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
