package graph
import (
    "strings"
    "errors"
)
// 
type DeviceManagementApplicabilityRuleType int

const (
    INCLUDE_DEVICEMANAGEMENTAPPLICABILITYRULETYPE DeviceManagementApplicabilityRuleType = iota
    EXCLUDE_DEVICEMANAGEMENTAPPLICABILITYRULETYPE
)

func (i DeviceManagementApplicabilityRuleType) String() string {
    return []string{"INCLUDE", "EXCLUDE"}[i]
}
func ParseDeviceManagementApplicabilityRuleType(v string) (interface{}, error) {
    result := INCLUDE_DEVICEMANAGEMENTAPPLICABILITYRULETYPE
    switch strings.ToUpper(v) {
        case "INCLUDE":
            result = INCLUDE_DEVICEMANAGEMENTAPPLICABILITYRULETYPE
        case "EXCLUDE":
            result = EXCLUDE_DEVICEMANAGEMENTAPPLICABILITYRULETYPE
        default:
            return 0, errors.New("Unknown DeviceManagementApplicabilityRuleType value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementApplicabilityRuleType(values []DeviceManagementApplicabilityRuleType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
