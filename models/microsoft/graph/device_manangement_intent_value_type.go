package graph
import (
    "strings"
    "errors"
)
// 
type DeviceManangementIntentValueType int

const (
    INTEGER_DEVICEMANANGEMENTINTENTVALUETYPE DeviceManangementIntentValueType = iota
    BOOLEAN_DEVICEMANANGEMENTINTENTVALUETYPE
    STRING_DEVICEMANANGEMENTINTENTVALUETYPE
    COMPLEX_DEVICEMANANGEMENTINTENTVALUETYPE
    COLLECTION_DEVICEMANANGEMENTINTENTVALUETYPE
    ABSTRACTCOMPLEX_DEVICEMANANGEMENTINTENTVALUETYPE
)

func (i DeviceManangementIntentValueType) String() string {
    return []string{"INTEGER", "BOOLEAN", "STRING", "COMPLEX", "COLLECTION", "ABSTRACTCOMPLEX"}[i]
}
func ParseDeviceManangementIntentValueType(v string) (interface{}, error) {
    result := INTEGER_DEVICEMANANGEMENTINTENTVALUETYPE
    switch strings.ToUpper(v) {
        case "INTEGER":
            result = INTEGER_DEVICEMANANGEMENTINTENTVALUETYPE
        case "BOOLEAN":
            result = BOOLEAN_DEVICEMANANGEMENTINTENTVALUETYPE
        case "STRING":
            result = STRING_DEVICEMANANGEMENTINTENTVALUETYPE
        case "COMPLEX":
            result = COMPLEX_DEVICEMANANGEMENTINTENTVALUETYPE
        case "COLLECTION":
            result = COLLECTION_DEVICEMANANGEMENTINTENTVALUETYPE
        case "ABSTRACTCOMPLEX":
            result = ABSTRACTCOMPLEX_DEVICEMANANGEMENTINTENTVALUETYPE
        default:
            return 0, errors.New("Unknown DeviceManangementIntentValueType value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManangementIntentValueType(values []DeviceManangementIntentValueType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
