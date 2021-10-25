package graph
import (
    "strings"
    "errors"
)
type DeviceManagementApplicabilityRuleType int

const (
    INCLUDE_DEVICEMANAGEMENTAPPLICABILITYRULETYPE DeviceManagementApplicabilityRuleType = iota
    EXCLUDE_DEVICEMANAGEMENTAPPLICABILITYRULETYPE
)

func (i DeviceManagementApplicabilityRuleType) String() string {
    return []string{"INCLUDE", "EXCLUDE"}[i]
}
func ParseDeviceManagementApplicabilityRuleType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "INCLUDE":
            return INCLUDE_DEVICEMANAGEMENTAPPLICABILITYRULETYPE, nil
        case "EXCLUDE":
            return EXCLUDE_DEVICEMANAGEMENTAPPLICABILITYRULETYPE, nil
    }
    return 0, errors.New("Unknown DeviceManagementApplicabilityRuleType value: " + v)
}
func SerializeDeviceManagementApplicabilityRuleType(values []DeviceManagementApplicabilityRuleType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
