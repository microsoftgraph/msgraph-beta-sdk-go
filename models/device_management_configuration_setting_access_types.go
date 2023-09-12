package models
import (
    "errors"
    "strings"
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
    var values []string
    for p := DeviceManagementConfigurationSettingAccessTypes(1); p <= EXECUTE_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"none", "add", "copy", "delete", "get", "replace", "execute"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseDeviceManagementConfigurationSettingAccessTypes(v string) (any, error) {
    var result DeviceManagementConfigurationSettingAccessTypes
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES
            case "add":
                result |= ADD_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES
            case "copy":
                result |= COPY_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES
            case "delete":
                result |= DELETE_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES
            case "get":
                result |= GET_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES
            case "replace":
                result |= REPLACE_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES
            case "execute":
                result |= EXECUTE_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES
            default:
                return 0, errors.New("Unknown DeviceManagementConfigurationSettingAccessTypes value: " + v)
        }
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
func (i DeviceManagementConfigurationSettingAccessTypes) isMultiValue() bool {
    return true
}
