package security
import (
    "errors"
    "math"
    "strings"
)
// 
type DeviceIdEntityIdentifier int

const (
    DEVICEID_DEVICEIDENTITYIDENTIFIER = 1
    UNKNOWNFUTUREVALUE_DEVICEIDENTITYIDENTIFIER = 2
)

func (i DeviceIdEntityIdentifier) String() string {
    var values []string
    options := []string{"deviceId", "unknownFutureValue"}
    for p := 0; p < 2; p++ {
        mantis := DeviceIdEntityIdentifier(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
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
