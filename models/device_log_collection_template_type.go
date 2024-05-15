package models
// Enum for the template type used for collecting logs
type DeviceLogCollectionTemplateType int

const (
    // Predefined template for what will be collected
    PREDEFINED_DEVICELOGCOLLECTIONTEMPLATETYPE DeviceLogCollectionTemplateType = iota
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_DEVICELOGCOLLECTIONTEMPLATETYPE
)

func (i DeviceLogCollectionTemplateType) String() string {
    return []string{"predefined", "unknownFutureValue"}[i]
}
func ParseDeviceLogCollectionTemplateType(v string) (any, error) {
    result := PREDEFINED_DEVICELOGCOLLECTIONTEMPLATETYPE
    switch v {
        case "predefined":
            result = PREDEFINED_DEVICELOGCOLLECTIONTEMPLATETYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DEVICELOGCOLLECTIONTEMPLATETYPE
        default:
            return nil, nil
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
func (i DeviceLogCollectionTemplateType) isMultiValue() bool {
    return false
}
