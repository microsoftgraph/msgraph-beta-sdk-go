package models
// Enum type used for DeviceActionCategory
type DeviceActionCategory int

const (
    // Action is performed on a single device alone
    SINGLE_DEVICEACTIONCATEGORY DeviceActionCategory = iota
    // Action is performed for a set of devices
    BULK_DEVICEACTIONCATEGORY
)

func (i DeviceActionCategory) String() string {
    return []string{"single", "bulk"}[i]
}
func ParseDeviceActionCategory(v string) (any, error) {
    result := SINGLE_DEVICEACTIONCATEGORY
    switch v {
        case "single":
            result = SINGLE_DEVICEACTIONCATEGORY
        case "bulk":
            result = BULK_DEVICEACTIONCATEGORY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDeviceActionCategory(values []DeviceActionCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DeviceActionCategory) isMultiValue() bool {
    return false
}
