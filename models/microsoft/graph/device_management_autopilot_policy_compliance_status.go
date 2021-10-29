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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_DEVICEMANAGEMENTAUTOPILOTPOLICYCOMPLIANCESTATUS, nil
        case "COMPLIANT":
            return COMPLIANT_DEVICEMANAGEMENTAUTOPILOTPOLICYCOMPLIANCESTATUS, nil
        case "INSTALLED":
            return INSTALLED_DEVICEMANAGEMENTAUTOPILOTPOLICYCOMPLIANCESTATUS, nil
        case "NOTCOMPLIANT":
            return NOTCOMPLIANT_DEVICEMANAGEMENTAUTOPILOTPOLICYCOMPLIANCESTATUS, nil
        case "NOTINSTALLED":
            return NOTINSTALLED_DEVICEMANAGEMENTAUTOPILOTPOLICYCOMPLIANCESTATUS, nil
        case "ERROR":
            return ERROR_DEVICEMANAGEMENTAUTOPILOTPOLICYCOMPLIANCESTATUS, nil
    }
    return 0, errors.New("Unknown DeviceManagementAutopilotPolicyComplianceStatus value: " + v)
}
func SerializeDeviceManagementAutopilotPolicyComplianceStatus(values []DeviceManagementAutopilotPolicyComplianceStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
