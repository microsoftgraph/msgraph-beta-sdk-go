package graph
import (
    "strings"
    "errors"
)
// 
type DeviceCustomAttributeValueType int

const (
    INTEGER_DEVICECUSTOMATTRIBUTEVALUETYPE DeviceCustomAttributeValueType = iota
    STRING_DEVICECUSTOMATTRIBUTEVALUETYPE
    DATETIME_DEVICECUSTOMATTRIBUTEVALUETYPE
)

func (i DeviceCustomAttributeValueType) String() string {
    return []string{"INTEGER", "STRING", "DATETIME"}[i]
}
func ParseDeviceCustomAttributeValueType(v string) (interface{}, error) {
    result := INTEGER_DEVICECUSTOMATTRIBUTEVALUETYPE
    switch strings.ToUpper(v) {
        case "INTEGER":
            result = INTEGER_DEVICECUSTOMATTRIBUTEVALUETYPE
        case "STRING":
            result = STRING_DEVICECUSTOMATTRIBUTEVALUETYPE
        case "DATETIME":
            result = DATETIME_DEVICECUSTOMATTRIBUTEVALUETYPE
        default:
            return 0, errors.New("Unknown DeviceCustomAttributeValueType value: " + v)
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
