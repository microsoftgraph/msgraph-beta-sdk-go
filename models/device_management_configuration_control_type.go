package models
import (
    "errors"
)
// Setting control type representation in the UX
type DeviceManagementConfigurationControlType int

const (
    // Default. UX uses default UX element base on setting type for the setting.
    DEFAULTESCAPED_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE DeviceManagementConfigurationControlType = iota
    // Display the setting in dropdown box.
    DROPDOWN_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
    // Display text input in small text input.
    SMALLTEXTBOX_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
    // Display text input in large text input.
    LARGETEXTBOX_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
    // Allow for toggle control type.
    TOGGLE_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
    // Allow for multiheader grid control type.
    MULTIHEADERGRID_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
    // Allow for context pane control type.
    CONTEXTPANE_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
)

func (i DeviceManagementConfigurationControlType) String() string {
    return []string{"default", "dropdown", "smallTextBox", "largeTextBox", "toggle", "multiheaderGrid", "contextPane", "unknownFutureValue"}[i]
}
func ParseDeviceManagementConfigurationControlType(v string) (any, error) {
    result := DEFAULTESCAPED_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
    switch v {
        case "default":
            result = DEFAULTESCAPED_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
        case "dropdown":
            result = DROPDOWN_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
        case "smallTextBox":
            result = SMALLTEXTBOX_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
        case "largeTextBox":
            result = LARGETEXTBOX_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
        case "toggle":
            result = TOGGLE_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
        case "multiheaderGrid":
            result = MULTIHEADERGRID_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
        case "contextPane":
            result = CONTEXTPANE_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
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
