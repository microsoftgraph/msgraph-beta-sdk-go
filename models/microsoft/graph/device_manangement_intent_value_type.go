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
    switch strings.ToUpper(v) {
        case "INTEGER":
            return INTEGER_DEVICEMANANGEMENTINTENTVALUETYPE, nil
        case "BOOLEAN":
            return BOOLEAN_DEVICEMANANGEMENTINTENTVALUETYPE, nil
        case "STRING":
            return STRING_DEVICEMANANGEMENTINTENTVALUETYPE, nil
        case "COMPLEX":
            return COMPLEX_DEVICEMANANGEMENTINTENTVALUETYPE, nil
        case "COLLECTION":
            return COLLECTION_DEVICEMANANGEMENTINTENTVALUETYPE, nil
        case "ABSTRACTCOMPLEX":
            return ABSTRACTCOMPLEX_DEVICEMANANGEMENTINTENTVALUETYPE, nil
    }
    return 0, errors.New("Unknown DeviceManangementIntentValueType value: " + v)
}
func SerializeDeviceManangementIntentValueType(values []DeviceManangementIntentValueType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
