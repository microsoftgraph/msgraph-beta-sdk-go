package graph
import (
    "strings"
    "errors"
)
// 
type DeviceManagementConfigurationTechnologies int

const (
    NONE_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES DeviceManagementConfigurationTechnologies = iota
    MDM_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
    WINDOWS10XMANAGEMENT_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
    CONFIGMANAGER_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
    MICROSOFTSENSE_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
    EXCHANGEONLINE_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
    LINUXMDM_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
    UNKNOWNFUTUREVALUE_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
)

func (i DeviceManagementConfigurationTechnologies) String() string {
    return []string{"NONE", "MDM", "WINDOWS10XMANAGEMENT", "CONFIGMANAGER", "MICROSOFTSENSE", "EXCHANGEONLINE", "LINUXMDM", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseDeviceManagementConfigurationTechnologies(v string) (interface{}, error) {
    result := NONE_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
        case "MDM":
            result = MDM_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
        case "WINDOWS10XMANAGEMENT":
            result = WINDOWS10XMANAGEMENT_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
        case "CONFIGMANAGER":
            result = CONFIGMANAGER_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
        case "MICROSOFTSENSE":
            result = MICROSOFTSENSE_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
        case "EXCHANGEONLINE":
            result = EXCHANGEONLINE_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
        case "LINUXMDM":
            result = LINUXMDM_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
        default:
            return 0, errors.New("Unknown DeviceManagementConfigurationTechnologies value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementConfigurationTechnologies(values []DeviceManagementConfigurationTechnologies) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
