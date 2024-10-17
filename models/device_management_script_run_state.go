package models
// Indicates the type of execution status of the device management script. This status provides insights into whether the script has been successfully executed, encountered errors, or is pending execution.
type DeviceManagementScriptRunState int

const (
    // Default value. Indicates that the script execution status is unknown for the device.
    UNKNOWN_DEVICEMANAGEMENTSCRIPTRUNSTATE DeviceManagementScriptRunState = iota
    // Indicates that the script ran successfully for the device.
    SUCCESS_DEVICEMANAGEMENTSCRIPTRUNSTATE
    // Indicates that the script resulted in failure for the device.
    FAIL_DEVICEMANAGEMENTSCRIPTRUNSTATE
    // Indicates that the discovery script was unable to complete on the device.
    SCRIPTERROR_DEVICEMANAGEMENTSCRIPTRUNSTATE
    // Indicates that the script has not yet started running on the device.
    PENDING_DEVICEMANAGEMENTSCRIPTRUNSTATE
    // Indicates that the script is not applicable for this device.
    NOTAPPLICABLE_DEVICEMANAGEMENTSCRIPTRUNSTATE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_DEVICEMANAGEMENTSCRIPTRUNSTATE
)

func (i DeviceManagementScriptRunState) String() string {
    return []string{"unknown", "success", "fail", "scriptError", "pending", "notApplicable", "unknownFutureValue"}[i]
}
func ParseDeviceManagementScriptRunState(v string) (any, error) {
    result := UNKNOWN_DEVICEMANAGEMENTSCRIPTRUNSTATE
    switch v {
        case "unknown":
            result = UNKNOWN_DEVICEMANAGEMENTSCRIPTRUNSTATE
        case "success":
            result = SUCCESS_DEVICEMANAGEMENTSCRIPTRUNSTATE
        case "fail":
            result = FAIL_DEVICEMANAGEMENTSCRIPTRUNSTATE
        case "scriptError":
            result = SCRIPTERROR_DEVICEMANAGEMENTSCRIPTRUNSTATE
        case "pending":
            result = PENDING_DEVICEMANAGEMENTSCRIPTRUNSTATE
        case "notApplicable":
            result = NOTAPPLICABLE_DEVICEMANAGEMENTSCRIPTRUNSTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DEVICEMANAGEMENTSCRIPTRUNSTATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDeviceManagementScriptRunState(values []DeviceManagementScriptRunState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DeviceManagementScriptRunState) isMultiValue() bool {
    return false
}
