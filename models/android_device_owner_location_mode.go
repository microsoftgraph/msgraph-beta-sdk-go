package models
import (
    "errors"
)
// Android Device Owner Location Mode Type
type AndroidDeviceOwnerLocationMode int

const (
    // No restrictions on the location setting and no specific behavior is set or enforced. This is the default
    NOTCONFIGURED_ANDROIDDEVICEOWNERLOCATIONMODE AndroidDeviceOwnerLocationMode = iota
    // Location setting is disabled on the device
    DISABLED_ANDROIDDEVICEOWNERLOCATIONMODE
    // Evolvable enumeration sentinel value. Do not use
    UNKNOWNFUTUREVALUE_ANDROIDDEVICEOWNERLOCATIONMODE
)

func (i AndroidDeviceOwnerLocationMode) String() string {
    return []string{"notConfigured", "disabled", "unknownFutureValue"}[i]
}
func ParseAndroidDeviceOwnerLocationMode(v string) (any, error) {
    result := NOTCONFIGURED_ANDROIDDEVICEOWNERLOCATIONMODE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_ANDROIDDEVICEOWNERLOCATIONMODE
        case "disabled":
            result = DISABLED_ANDROIDDEVICEOWNERLOCATIONMODE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ANDROIDDEVICEOWNERLOCATIONMODE
        default:
            return 0, errors.New("Unknown AndroidDeviceOwnerLocationMode value: " + v)
    }
    return &result, nil
}
func SerializeAndroidDeviceOwnerLocationMode(values []AndroidDeviceOwnerLocationMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AndroidDeviceOwnerLocationMode) isMultiValue() bool {
    return false
}
