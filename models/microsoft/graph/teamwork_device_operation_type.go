package graph
import (
    "strings"
    "errors"
)
// 
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
    return []string{"DEVICERESTART", "CONFIGUPDATE", "DEVICEDIAGNOSTICS", "SOFTWAREUPDATE", "DEVICEMANAGEMENTAGENTCONFIGUPDATE", "REMOTELOGIN", "REMOTELOGOUT", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseTeamworkDeviceOperationType(v string) (interface{}, error) {
    result := DEVICERESTART_TEAMWORKDEVICEOPERATIONTYPE
    switch strings.ToUpper(v) {
        case "DEVICERESTART":
            result = DEVICERESTART_TEAMWORKDEVICEOPERATIONTYPE
        case "CONFIGUPDATE":
            result = CONFIGUPDATE_TEAMWORKDEVICEOPERATIONTYPE
        case "DEVICEDIAGNOSTICS":
            result = DEVICEDIAGNOSTICS_TEAMWORKDEVICEOPERATIONTYPE
        case "SOFTWAREUPDATE":
            result = SOFTWAREUPDATE_TEAMWORKDEVICEOPERATIONTYPE
        case "DEVICEMANAGEMENTAGENTCONFIGUPDATE":
            result = DEVICEMANAGEMENTAGENTCONFIGUPDATE_TEAMWORKDEVICEOPERATIONTYPE
        case "REMOTELOGIN":
            result = REMOTELOGIN_TEAMWORKDEVICEOPERATIONTYPE
        case "REMOTELOGOUT":
            result = REMOTELOGOUT_TEAMWORKDEVICEOPERATIONTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_TEAMWORKDEVICEOPERATIONTYPE
        default:
            return 0, errors.New("Unknown TeamworkDeviceOperationType value: " + v)
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
