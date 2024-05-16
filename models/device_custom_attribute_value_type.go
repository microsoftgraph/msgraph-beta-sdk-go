package models
// Represents the expected type for a macOS custom attribute script value.
type DeviceCustomAttributeValueType int

const (
    // Indicates the value for a custom attribute script is an integer.
    INTEGER_DEVICECUSTOMATTRIBUTEVALUETYPE DeviceCustomAttributeValueType = iota
    // Indicates the value for a custom attribute script is a string.
    STRING_DEVICECUSTOMATTRIBUTEVALUETYPE
    // Indicates the value for a custom attribute script is a date conforming to ISO 8601.
    DATETIME_DEVICECUSTOMATTRIBUTEVALUETYPE
)

func (i DeviceCustomAttributeValueType) String() string {
    return []string{"integer", "string", "dateTime"}[i]
}
func ParseDeviceCustomAttributeValueType(v string) (any, error) {
    result := INTEGER_DEVICECUSTOMATTRIBUTEVALUETYPE
    switch v {
        case "integer":
            result = INTEGER_DEVICECUSTOMATTRIBUTEVALUETYPE
        case "string":
            result = STRING_DEVICECUSTOMATTRIBUTEVALUETYPE
        case "dateTime":
            result = DATETIME_DEVICECUSTOMATTRIBUTEVALUETYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDeviceCustomAttributeValueType(values []DeviceCustomAttributeValueType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DeviceCustomAttributeValueType) isMultiValue() bool {
    return false
}
