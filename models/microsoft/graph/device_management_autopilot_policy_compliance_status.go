package graph
import (
    "strings"
    "errors"
)
// 
type DeviceManagementAutopilotPolicyComplianceStatus int

const (
    UNKNOWN_DEVICEMANAGEMENTAUTOPILOTPOLICYCOMPLIANCESTATUS DeviceManagementAutopilotPolicyComplianceStatus = iota
    COMPLIANT_DEVICEMANAGEMENTAUTOPILOTPOLICYCOMPLIANCESTATUS
    INSTALLED_DEVICEMANAGEMENTAUTOPILOTPOLICYCOMPLIANCESTATUS
    NOTCOMPLIANT_DEVICEMANAGEMENTAUTOPILOTPOLICYCOMPLIANCESTATUS
    NOTINSTALLED_DEVICEMANAGEMENTAUTOPILOTPOLICYCOMPLIANCESTATUS
    ERROR_DEVICEMANAGEMENTAUTOPILOTPOLICYCOMPLIANCESTATUS
)

func (i DeviceManagementAutopilotPolicyComplianceStatus) String() string {
    return []string{"UNKNOWN", "COMPLIANT", "INSTALLED", "NOTCOMPLIANT", "NOTINSTALLED", "ERROR"}[i]
}
func ParseDeviceManagementAutopilotPolicyComplianceStatus(v string) (interface{}, error) {
    result := UNKNOWN_DEVICEMANAGEMENTAUTOPILOTPOLICYCOMPLIANCESTATUS
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_DEVICEMANAGEMENTAUTOPILOTPOLICYCOMPLIANCESTATUS
        case "COMPLIANT":
            result = COMPLIANT_DEVICEMANAGEMENTAUTOPILOTPOLICYCOMPLIANCESTATUS
        case "INSTALLED":
            result = INSTALLED_DEVICEMANAGEMENTAUTOPILOTPOLICYCOMPLIANCESTATUS
        case "NOTCOMPLIANT":
            result = NOTCOMPLIANT_DEVICEMANAGEMENTAUTOPILOTPOLICYCOMPLIANCESTATUS
        case "NOTINSTALLED":
            result = NOTINSTALLED_DEVICEMANAGEMENTAUTOPILOTPOLICYCOMPLIANCESTATUS
        case "ERROR":
            result = ERROR_DEVICEMANAGEMENTAUTOPILOTPOLICYCOMPLIANCESTATUS
        default:
            return 0, errors.New("Unknown DeviceManagementAutopilotPolicyComplianceStatus value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementAutopilotPolicyComplianceStatus(values []DeviceManagementAutopilotPolicyComplianceStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
