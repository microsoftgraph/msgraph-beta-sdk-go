package graph
import (
    "strings"
    "errors"
)
type DeviceManagementConfigurationDeviceMode int

const (
    NONE_DEVICEMANAGEMENTCONFIGURATIONDEVICEMODE DeviceManagementConfigurationDeviceMode = iota
    KIOSK_DEVICEMANAGEMENTCONFIGURATIONDEVICEMODE
)

func (i DeviceManagementConfigurationDeviceMode) String() string {
    return []string{"NONE", "KIOSK"}[i]
}
func ParseDeviceManagementConfigurationDeviceMode(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_DEVICEMANAGEMENTCONFIGURATIONDEVICEMODE, nil
        case "KIOSK":
            return KIOSK_DEVICEMANAGEMENTCONFIGURATIONDEVICEMODE, nil
    }
    return 0, errors.New("Unknown DeviceManagementConfigurationDeviceMode value: " + v)
}
func SerializeDeviceManagementConfigurationDeviceMode(values []DeviceManagementConfigurationDeviceMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
