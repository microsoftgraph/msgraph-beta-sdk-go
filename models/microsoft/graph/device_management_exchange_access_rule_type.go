package graph
import (
    "strings"
    "errors"
)
// 
type DeviceManagementExchangeAccessRuleType int

const (
    FAMILY_DEVICEMANAGEMENTEXCHANGEACCESSRULETYPE DeviceManagementExchangeAccessRuleType = iota
    MODEL_DEVICEMANAGEMENTEXCHANGEACCESSRULETYPE
)

func (i DeviceManagementExchangeAccessRuleType) String() string {
    return []string{"FAMILY", "MODEL"}[i]
}
func ParseDeviceManagementExchangeAccessRuleType(v string) (interface{}, error) {
    result := FAMILY_DEVICEMANAGEMENTEXCHANGEACCESSRULETYPE
    switch strings.ToUpper(v) {
        case "FAMILY":
            result = FAMILY_DEVICEMANAGEMENTEXCHANGEACCESSRULETYPE
        case "MODEL":
            result = MODEL_DEVICEMANAGEMENTEXCHANGEACCESSRULETYPE
        default:
            return 0, errors.New("Unknown DeviceManagementExchangeAccessRuleType value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementExchangeAccessRuleType(values []DeviceManagementExchangeAccessRuleType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
