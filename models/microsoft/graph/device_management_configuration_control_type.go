package graph
import (
    "strings"
    "errors"
)
// 
type DeviceManagementConfigurationControlType int

const (
    DEFAULT_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE DeviceManagementConfigurationControlType = iota
    DROPDOWN_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
    SMALLTEXTBOX_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
    LARGETEXTBOX_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
    TOGGLE_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
    MULTIHEADERGRID_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
    CONTEXTPANE_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
)

func (i DeviceManagementConfigurationControlType) String() string {
    return []string{"DEFAULT", "DROPDOWN", "SMALLTEXTBOX", "LARGETEXTBOX", "TOGGLE", "MULTIHEADERGRID", "CONTEXTPANE"}[i]
}
func ParseDeviceManagementConfigurationControlType(v string) (interface{}, error) {
    result := DEFAULT_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
    switch strings.ToUpper(v) {
        case "DEFAULT":
            result = DEFAULT_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
        case "DROPDOWN":
            result = DROPDOWN_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
        case "SMALLTEXTBOX":
            result = SMALLTEXTBOX_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
        case "LARGETEXTBOX":
            result = LARGETEXTBOX_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
        case "TOGGLE":
            result = TOGGLE_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
        case "MULTIHEADERGRID":
            result = MULTIHEADERGRID_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
        case "CONTEXTPANE":
            result = CONTEXTPANE_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
        default:
            return 0, errors.New("Unknown DeviceManagementConfigurationControlType value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementConfigurationControlType(values []DeviceManagementConfigurationControlType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
