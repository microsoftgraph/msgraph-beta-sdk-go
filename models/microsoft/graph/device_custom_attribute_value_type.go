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
    switch strings.ToUpper(v) {
        case "INTEGER":
            return INTEGER_DEVICECUSTOMATTRIBUTEVALUETYPE, nil
        case "STRING":
            return STRING_DEVICECUSTOMATTRIBUTEVALUETYPE, nil
        case "DATETIME":
            return DATETIME_DEVICECUSTOMATTRIBUTEVALUETYPE, nil
    }
    return 0, errors.New("Unknown DeviceCustomAttributeValueType value: " + v)
}
func SerializeDeviceCustomAttributeValueType(values []DeviceCustomAttributeValueType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
