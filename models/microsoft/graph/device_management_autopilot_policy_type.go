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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_DEVICEMANAGEMENTAUTOPILOTPOLICYTYPE, nil
        case "APPLICATION":
            return APPLICATION_DEVICEMANAGEMENTAUTOPILOTPOLICYTYPE, nil
        case "APPMODEL":
            return APPMODEL_DEVICEMANAGEMENTAUTOPILOTPOLICYTYPE, nil
        case "CONFIGURATIONPOLICY":
            return CONFIGURATIONPOLICY_DEVICEMANAGEMENTAUTOPILOTPOLICYTYPE, nil
    }
    return 0, errors.New("Unknown DeviceManagementAutopilotPolicyType value: " + v)
}
func SerializeDeviceManagementAutopilotPolicyType(values []DeviceManagementAutopilotPolicyType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
