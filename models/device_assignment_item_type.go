package models
// A list of possible device assignment item types to execute this action on the managed device. Device assignment item represents existing assigned Intune resource such as application or configuration. Currently supported device assignment item types are Application, DeviceConfiguration, DeviceManagementConfigurationPolicy and MobileAppConfiguration
type DeviceAssignmentItemType int

const (
    // Default. Indicates that the device assignment item type for the action is graph.mobileApp. Application is uninstalled on removal and installed back on restoration
    APPLICATION_DEVICEASSIGNMENTITEMTYPE DeviceAssignmentItemType = iota
    // Indicates that the device assignment item type for the action is graph.deviceConfiguration. DeviceConfiguration associated settings are removed on removal and added back on restoration
    DEVICECONFIGURATION_DEVICEASSIGNMENTITEMTYPE
    // Indicates that the device assignment item type for the action is graph.deviceManagementConfigurationPolicy. DeviceManagementConfigurationPolicy associated settings are removed on removal and added back on restoration
    DEVICEMANAGEMENTCONFIGURATIONPOLICY_DEVICEASSIGNMENTITEMTYPE
    // Indicates that the device assignment item type for the action is `graph.managedDeviceMobileAppConfiguration`. MobileAppConfiguration associated settings are removed on removal and added back on restoration
    MOBILEAPPCONFIGURATION_DEVICEASSIGNMENTITEMTYPE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_DEVICEASSIGNMENTITEMTYPE
)

func (i DeviceAssignmentItemType) String() string {
    return []string{"application", "deviceConfiguration", "deviceManagementConfigurationPolicy", "mobileAppConfiguration", "unknownFutureValue"}[i]
}
func ParseDeviceAssignmentItemType(v string) (any, error) {
    result := APPLICATION_DEVICEASSIGNMENTITEMTYPE
    switch v {
        case "application":
            result = APPLICATION_DEVICEASSIGNMENTITEMTYPE
        case "deviceConfiguration":
            result = DEVICECONFIGURATION_DEVICEASSIGNMENTITEMTYPE
        case "deviceManagementConfigurationPolicy":
            result = DEVICEMANAGEMENTCONFIGURATIONPOLICY_DEVICEASSIGNMENTITEMTYPE
        case "mobileAppConfiguration":
            result = MOBILEAPPCONFIGURATION_DEVICEASSIGNMENTITEMTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DEVICEASSIGNMENTITEMTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDeviceAssignmentItemType(values []DeviceAssignmentItemType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DeviceAssignmentItemType) isMultiValue() bool {
    return false
}
