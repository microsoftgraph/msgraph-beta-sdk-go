// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models
type TeamworkDeviceOperationType int

const (
    DEVICERESTART_TEAMWORKDEVICEOPERATIONTYPE TeamworkDeviceOperationType = iota
    CONFIGUPDATE_TEAMWORKDEVICEOPERATIONTYPE
    DEVICEDIAGNOSTICS_TEAMWORKDEVICEOPERATIONTYPE
    SOFTWAREUPDATE_TEAMWORKDEVICEOPERATIONTYPE
    DEVICEMANAGEMENTAGENTCONFIGUPDATE_TEAMWORKDEVICEOPERATIONTYPE
    REMOTELOGIN_TEAMWORKDEVICEOPERATIONTYPE
    REMOTELOGOUT_TEAMWORKDEVICEOPERATIONTYPE
    UNKNOWNFUTUREVALUE_TEAMWORKDEVICEOPERATIONTYPE
)

func (i TeamworkDeviceOperationType) String() string {
    return []string{"deviceRestart", "configUpdate", "deviceDiagnostics", "softwareUpdate", "deviceManagementAgentConfigUpdate", "remoteLogin", "remoteLogout", "unknownFutureValue"}[i]
}
func ParseTeamworkDeviceOperationType(v string) (any, error) {
    result := DEVICERESTART_TEAMWORKDEVICEOPERATIONTYPE
    switch v {
        case "deviceRestart":
            result = DEVICERESTART_TEAMWORKDEVICEOPERATIONTYPE
        case "configUpdate":
            result = CONFIGUPDATE_TEAMWORKDEVICEOPERATIONTYPE
        case "deviceDiagnostics":
            result = DEVICEDIAGNOSTICS_TEAMWORKDEVICEOPERATIONTYPE
        case "softwareUpdate":
            result = SOFTWAREUPDATE_TEAMWORKDEVICEOPERATIONTYPE
        case "deviceManagementAgentConfigUpdate":
            result = DEVICEMANAGEMENTAGENTCONFIGUPDATE_TEAMWORKDEVICEOPERATIONTYPE
        case "remoteLogin":
            result = REMOTELOGIN_TEAMWORKDEVICEOPERATIONTYPE
        case "remoteLogout":
            result = REMOTELOGOUT_TEAMWORKDEVICEOPERATIONTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TEAMWORKDEVICEOPERATIONTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeTeamworkDeviceOperationType(values []TeamworkDeviceOperationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TeamworkDeviceOperationType) isMultiValue() bool {
    return false
}
