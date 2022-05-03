package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type DeviceScopeOperator int

const (
    NONE_DEVICESCOPEOPERATOR DeviceScopeOperator = iota
    EQUALS_DEVICESCOPEOPERATOR
    UNKNOWNFUTUREVALUE_DEVICESCOPEOPERATOR
)

func (i DeviceScopeOperator) String() string {
    return []string{"NONE", "EQUALS", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseDeviceScopeOperator(v string) (interface{}, error) {
    result := NONE_DEVICESCOPEOPERATOR
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_DEVICESCOPEOPERATOR
        case "EQUALS":
            result = EQUALS_DEVICESCOPEOPERATOR
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_DEVICESCOPEOPERATOR
        default:
            return 0, errors.New("Unknown DeviceScopeOperator value: " + v)
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
