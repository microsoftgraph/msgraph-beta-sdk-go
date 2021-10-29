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
    switch strings.ToUpper(v) {
        case "FAMILY":
            return FAMILY_DEVICEMANAGEMENTEXCHANGEACCESSRULETYPE, nil
        case "MODEL":
            return MODEL_DEVICEMANAGEMENTEXCHANGEACCESSRULETYPE, nil
    }
    return 0, errors.New("Unknown DeviceManagementExchangeAccessRuleType value: " + v)
}
func SerializeDeviceManagementExchangeAccessRuleType(values []DeviceManagementExchangeAccessRuleType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
