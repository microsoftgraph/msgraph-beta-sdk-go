package models
import (
    "errors"
    "math"
    "strings"
)
type DeviceManagementConfigurationSettingAccessTypes int

const (
    NONE_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES = 1
    ADD_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES = 2
    COPY_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES = 4
    DELETE_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES = 8
    GET_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES = 16
    REPLACE_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES = 32
    EXECUTE_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES = 64
)

func (i DeviceManagementConfigurationSettingAccessTypes) String() string {
    var values []string
    options := []string{"none", "add", "copy", "delete", "get", "replace", "execute"}
    for p := 0; p < 7; p++ {
        mantis := DeviceManagementConfigurationSettingAccessTypes(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
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
