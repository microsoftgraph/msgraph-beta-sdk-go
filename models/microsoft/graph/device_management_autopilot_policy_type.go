package graph
import (
    "strings"
    "errors"
)
// 
type DeviceManagementAutopilotPolicyType int

const (
    UNKNOWN_DEVICEMANAGEMENTAUTOPILOTPOLICYTYPE DeviceManagementAutopilotPolicyType = iota
    APPLICATION_DEVICEMANAGEMENTAUTOPILOTPOLICYTYPE
    APPMODEL_DEVICEMANAGEMENTAUTOPILOTPOLICYTYPE
    CONFIGURATIONPOLICY_DEVICEMANAGEMENTAUTOPILOTPOLICYTYPE
)

func (i DeviceManagementAutopilotPolicyType) String() string {
    return []string{"UNKNOWN", "APPLICATION", "APPMODEL", "CONFIGURATIONPOLICY"}[i]
}
func ParseDeviceManagementAutopilotPolicyType(v string) (interface{}, error) {
    result := UNKNOWN_DEVICEMANAGEMENTAUTOPILOTPOLICYTYPE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_DEVICEMANAGEMENTAUTOPILOTPOLICYTYPE
        case "APPLICATION":
            result = APPLICATION_DEVICEMANAGEMENTAUTOPILOTPOLICYTYPE
        case "APPMODEL":
            result = APPMODEL_DEVICEMANAGEMENTAUTOPILOTPOLICYTYPE
        case "CONFIGURATIONPOLICY":
            result = CONFIGURATIONPOLICY_DEVICEMANAGEMENTAUTOPILOTPOLICYTYPE
        default:
            return 0, errors.New("Unknown DeviceManagementAutopilotPolicyType value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementAutopilotPolicyType(values []DeviceManagementAutopilotPolicyType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
