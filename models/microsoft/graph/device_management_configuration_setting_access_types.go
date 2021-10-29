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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES, nil
        case "ADD":
            return ADD_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES, nil
        case "COPY":
            return COPY_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES, nil
        case "DELETE":
            return DELETE_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES, nil
        case "GET":
            return GET_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES, nil
        case "REPLACE":
            return REPLACE_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES, nil
        case "EXECUTE":
            return EXECUTE_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES, nil
    }
    return 0, errors.New("Unknown DeviceManagementConfigurationSettingAccessTypes value: " + v)
}
func SerializeDeviceManagementConfigurationSettingAccessTypes(values []DeviceManagementConfigurationSettingAccessTypes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
