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
    switch strings.ToUpper(v) {
        case "DEVICERESTART":
            return DEVICERESTART_TEAMWORKDEVICEOPERATIONTYPE, nil
        case "CONFIGUPDATE":
            return CONFIGUPDATE_TEAMWORKDEVICEOPERATIONTYPE, nil
        case "DEVICEDIAGNOSTICS":
            return DEVICEDIAGNOSTICS_TEAMWORKDEVICEOPERATIONTYPE, nil
        case "SOFTWAREUPDATE":
            return SOFTWAREUPDATE_TEAMWORKDEVICEOPERATIONTYPE, nil
        case "DEVICEMANAGEMENTAGENTCONFIGUPDATE":
            return DEVICEMANAGEMENTAGENTCONFIGUPDATE_TEAMWORKDEVICEOPERATIONTYPE, nil
        case "REMOTELOGIN":
            return REMOTELOGIN_TEAMWORKDEVICEOPERATIONTYPE, nil
        case "REMOTELOGOUT":
            return REMOTELOGOUT_TEAMWORKDEVICEOPERATIONTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_TEAMWORKDEVICEOPERATIONTYPE, nil
    }
    return 0, errors.New("Unknown TeamworkDeviceOperationType value: " + v)
}
func SerializeTeamworkDeviceOperationType(values []TeamworkDeviceOperationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
