package graph
import (
    "strings"
    "errors"
)
// 
type DeviceManagementConfigurationSettingAccessTypes int

const (
    NONE_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES DeviceManagementConfigurationSettingAccessTypes = iota
    ADD_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES
    COPY_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES
    DELETE_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES
    GET_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES
    REPLACE_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES
    EXECUTE_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES
)

func (i DeviceManagementConfigurationSettingAccessTypes) String() string {
    return []string{"NONE", "ADD", "COPY", "DELETE", "GET", "REPLACE", "EXECUTE"}[i]
}
func ParseDeviceManagementConfigurationSettingAccessTypes(v string) (interface{}, error) {
    result := NONE_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES
        case "ADD":
            result = ADD_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES
        case "COPY":
            result = COPY_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES
        case "DELETE":
            result = DELETE_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES
        case "GET":
            result = GET_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES
        case "REPLACE":
            result = REPLACE_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES
        case "EXECUTE":
            result = EXECUTE_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES
        default:
            return 0, errors.New("Unknown DeviceManagementConfigurationSettingAccessTypes value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementConfigurationSettingAccessTypes(values []DeviceManagementConfigurationSettingAccessTypes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
