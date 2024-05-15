package models
import (
    "math"
    "strings"
)
// Describes which technology this setting can be deployed with
type DeviceManagementConfigurationTechnologies int

const (
    // Default. Setting cannot be deployed through any channel.
    NONE_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES = 1
    // Setting can be deployed through the MDM channel.
    MDM_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES = 2
    // Setting can be deployed through the Windows10XManagement channel
    WINDOWS10XMANAGEMENT_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES = 4
    // Setting can be deployed through the ConfigManager channel.
    CONFIGMANAGER_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES = 8
    // Setting can be deployed through the AppleRemoteManagement channel.
    APPLEREMOTEMANAGEMENT_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES = 16
    // Setting can be deployed through the SENSE agent channel.
    MICROSOFTSENSE_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES = 32
    // Setting can be deployed through the Exchange Online agent channel.
    EXCHANGEONLINE_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES = 64
    // Setting can be deployed through the Mobile Application Management (MAM) channel
    MOBILEAPPLICATIONMANAGEMENT_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES = 128
    // Setting can be deployed through the Linux Mdm channel.
    LINUXMDM_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES = 256
    // Setting can be deployed through device enrollment.
    ENROLLMENT_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES = 512
    // Setting can be deployed using the Endpoint privilege management channel
    ENDPOINTPRIVILEGEMANAGEMENT_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES = 1024
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES = 2048
    // Setting can be deployed using the Operating System Recovery channel
    WINDOWSOSRECOVERY_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES = 4096
)

func (i DeviceManagementConfigurationTechnologies) String() string {
    var values []string
    options := []string{"none", "mdm", "windows10XManagement", "configManager", "appleRemoteManagement", "microsoftSense", "exchangeOnline", "mobileApplicationManagement", "linuxMdm", "enrollment", "endpointPrivilegeManagement", "unknownFutureValue", "windowsOsRecovery"}
    for p := 0; p < 13; p++ {
        mantis := DeviceManagementConfigurationTechnologies(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseDeviceManagementConfigurationTechnologies(v string) (any, error) {
    var result DeviceManagementConfigurationTechnologies
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
            case "mdm":
                result |= MDM_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
            case "windows10XManagement":
                result |= WINDOWS10XMANAGEMENT_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
            case "configManager":
                result |= CONFIGMANAGER_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
            case "appleRemoteManagement":
                result |= APPLEREMOTEMANAGEMENT_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
            case "microsoftSense":
                result |= MICROSOFTSENSE_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
            case "exchangeOnline":
                result |= EXCHANGEONLINE_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
            case "mobileApplicationManagement":
                result |= MOBILEAPPLICATIONMANAGEMENT_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
            case "linuxMdm":
                result |= LINUXMDM_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
            case "enrollment":
                result |= ENROLLMENT_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
            case "endpointPrivilegeManagement":
                result |= ENDPOINTPRIVILEGEMANAGEMENT_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
            case "windowsOsRecovery":
                result |= WINDOWSOSRECOVERY_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
            default:
                return nil, nil
        }
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
func (i DeviceManagementConfigurationTechnologies) isMultiValue() bool {
    return true
}
