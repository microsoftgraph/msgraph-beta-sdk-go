package models
// Device scope configuration query operator. Possible values are: equals, notEquals, contains, notContains, greaterThan, lessThan. Default value: equals.
type DeviceScopeOperator int

const (
    // No operator set for the device scope configuration.
    NONE_DEVICESCOPEOPERATOR DeviceScopeOperator = iota
    // Operator for the device configuration query to be used (Equals).
    EQUALS_DEVICESCOPEOPERATOR
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_DEVICESCOPEOPERATOR
)

func (i DeviceScopeOperator) String() string {
    return []string{"none", "equals", "unknownFutureValue"}[i]
}
func ParseDeviceScopeOperator(v string) (any, error) {
    result := NONE_DEVICESCOPEOPERATOR
    switch v {
        case "none":
            result = NONE_DEVICESCOPEOPERATOR
        case "equals":
            result = EQUALS_DEVICESCOPEOPERATOR
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DEVICESCOPEOPERATOR
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDeviceScopeOperator(values []DeviceScopeOperator) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DeviceScopeOperator) isMultiValue() bool {
    return false
}
