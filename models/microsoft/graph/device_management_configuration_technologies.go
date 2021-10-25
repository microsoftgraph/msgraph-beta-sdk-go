package graph
import (
    "strings"
    "errors"
)
type DeviceManagementConfigurationTechnologies int

const (
    NONE_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES DeviceManagementConfigurationTechnologies = iota
    MDM_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
    WINDOWS10XMANAGEMENT_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
    CONFIGMANAGER_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
    MICROSOFTSENSE_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
)

func (i DeviceManagementConfigurationTechnologies) String() string {
    return []string{"NONE", "MDM", "WINDOWS10XMANAGEMENT", "CONFIGMANAGER", "MICROSOFTSENSE"}[i]
}
func ParseDeviceManagementConfigurationTechnologies(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES, nil
        case "MDM":
            return MDM_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES, nil
        case "WINDOWS10XMANAGEMENT":
            return WINDOWS10XMANAGEMENT_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES, nil
        case "CONFIGMANAGER":
            return CONFIGMANAGER_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES, nil
        case "MICROSOFTSENSE":
            return MICROSOFTSENSE_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES, nil
    }
    return 0, errors.New("Unknown DeviceManagementConfigurationTechnologies value: " + v)
}
func SerializeDeviceManagementConfigurationTechnologies(values []DeviceManagementConfigurationTechnologies) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
