package security
import (
    "errors"
    "strings"
)
// 
type DeviceIdEntityIdentifier int

const (
    DEVICEID_DEVICEIDENTITYIDENTIFIER DeviceIdEntityIdentifier = iota
    UNKNOWNFUTUREVALUE_DEVICEIDENTITYIDENTIFIER
)

func (i DeviceIdEntityIdentifier) String() string {
    var values []string
    for p := DeviceIdEntityIdentifier(1); p <= UNKNOWNFUTUREVALUE_DEVICEIDENTITYIDENTIFIER; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"deviceId", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseDeviceIdEntityIdentifier(v string) (any, error) {
    var result DeviceIdEntityIdentifier
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "deviceId":
                result |= DEVICEID_DEVICEIDENTITYIDENTIFIER
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_DEVICEIDENTITYIDENTIFIER
            default:
                return 0, errors.New("Unknown DeviceIdEntityIdentifier value: " + v)
        }
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
func (i DeviceIdEntityIdentifier) isMultiValue() bool {
    return true
}
