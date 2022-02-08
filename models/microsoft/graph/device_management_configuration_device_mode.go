package graph
import (
    "strings"
    "errors"
)
// 
type DeviceManagementConfigurationDeviceMode int

const (
    NONE_DEVICEMANAGEMENTCONFIGURATIONDEVICEMODE DeviceManagementConfigurationDeviceMode = iota
    KIOSK_DEVICEMANAGEMENTCONFIGURATIONDEVICEMODE
)

func (i DeviceManagementConfigurationDeviceMode) String() string {
    return []string{"NONE", "KIOSK"}[i]
}
func ParseDeviceManagementConfigurationDeviceMode(v string) (interface{}, error) {
    result := NONE_DEVICEMANAGEMENTCONFIGURATIONDEVICEMODE
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_DEVICEMANAGEMENTCONFIGURATIONDEVICEMODE
        case "KIOSK":
            result = KIOSK_DEVICEMANAGEMENTCONFIGURATIONDEVICEMODE
        default:
            return 0, errors.New("Unknown DeviceManagementConfigurationDeviceMode value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementConfigurationDeviceMode(values []DeviceManagementConfigurationDeviceMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
