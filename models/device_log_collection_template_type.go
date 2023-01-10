package models
import (
    "errors"
)
// Enum for the template type used for collecting logs
type DeviceLogCollectionTemplateType int

const (
    // Predefined template for what will be collected
    PREDEFINED_DEVICELOGCOLLECTIONTEMPLATETYPE DeviceLogCollectionTemplateType = iota
)

func (i DeviceLogCollectionTemplateType) String() string {
    return []string{"predefined"}[i]
}
func ParseDeviceLogCollectionTemplateType(v string) (interface{}, error) {
    result := PREDEFINED_DEVICELOGCOLLECTIONTEMPLATETYPE
    switch v {
        case "predefined":
            result = PREDEFINED_DEVICELOGCOLLECTIONTEMPLATETYPE
        default:
            return 0, errors.New("Unknown DeviceLogCollectionTemplateType value: " + v)
    }
    return &result, nil
}
func SerializeDeviceLogCollectionTemplateType(values []DeviceLogCollectionTemplateType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
