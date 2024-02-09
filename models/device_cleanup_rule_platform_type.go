package models
import (
    "errors"
)
// Define the platform type for which the admin wants to create the device clean up rule
type DeviceCleanupRulePlatformType int

const (
    // Default. Indicates that clean up rule is associated with all managed device platforms.
    ALL_DEVICECLEANUPRULEPLATFORMTYPE DeviceCleanupRulePlatformType = iota
    // Indicates that clean up rule is associated with Android open source project managed device platforms.
    ANDROIDAOSP_DEVICECLEANUPRULEPLATFORMTYPE
    // Indicates that clean up rule is associated with Android device administrator managed device platforms.
    ANDROIDDEVICEADMINISTRATOR_DEVICECLEANUPRULEPLATFORMTYPE
    // Indicates that clean up rule is associated with Android dedicated and fully managed and Corporate Owned Work Profile managed device platforms.
    ANDROIDDEDICATEDANDFULLYMANAGEDCORPORATEOWNEDWORKPROFILE_DEVICECLEANUPRULEPLATFORMTYPE
    // Indicates that clean up rule is associated with ChromeOS managed device platforms.
    CHROMEOS_DEVICECLEANUPRULEPLATFORMTYPE
    // Indicates that clean up rule is associated with Android personally owned work profile managed device platforms.
    ANDROIDPERSONALLYOWNEDWORKPROFILE_DEVICECLEANUPRULEPLATFORMTYPE
    // Indicates that clean up rule is associated with IOS managed device platforms.
    IOS_DEVICECLEANUPRULEPLATFORMTYPE
    // Indicates that clean up rule is associated with MacOS managed device platforms.
    MACOS_DEVICECLEANUPRULEPLATFORMTYPE
    // Indicates that clean up rule is associated with Windows managed device platforms.
    WINDOWS_DEVICECLEANUPRULEPLATFORMTYPE
    // Indicates that clean up rule is associated with Windows Holographic managed device platforms.
    WINDOWSHOLOGRAPHIC_DEVICECLEANUPRULEPLATFORMTYPE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_DEVICECLEANUPRULEPLATFORMTYPE
)

func (i DeviceCleanupRulePlatformType) String() string {
    return []string{"all", "androidAOSP", "androidDeviceAdministrator", "androidDedicatedAndFullyManagedCorporateOwnedWorkProfile", "chromeOS", "androidPersonallyOwnedWorkProfile", "ios", "macOS", "windows", "windowsHolographic", "unknownFutureValue"}[i]
}
func ParseDeviceCleanupRulePlatformType(v string) (any, error) {
    result := ALL_DEVICECLEANUPRULEPLATFORMTYPE
    switch v {
        case "all":
            result = ALL_DEVICECLEANUPRULEPLATFORMTYPE
        case "androidAOSP":
            result = ANDROIDAOSP_DEVICECLEANUPRULEPLATFORMTYPE
        case "androidDeviceAdministrator":
            result = ANDROIDDEVICEADMINISTRATOR_DEVICECLEANUPRULEPLATFORMTYPE
        case "androidDedicatedAndFullyManagedCorporateOwnedWorkProfile":
            result = ANDROIDDEDICATEDANDFULLYMANAGEDCORPORATEOWNEDWORKPROFILE_DEVICECLEANUPRULEPLATFORMTYPE
        case "chromeOS":
            result = CHROMEOS_DEVICECLEANUPRULEPLATFORMTYPE
        case "androidPersonallyOwnedWorkProfile":
            result = ANDROIDPERSONALLYOWNEDWORKPROFILE_DEVICECLEANUPRULEPLATFORMTYPE
        case "ios":
            result = IOS_DEVICECLEANUPRULEPLATFORMTYPE
        case "macOS":
            result = MACOS_DEVICECLEANUPRULEPLATFORMTYPE
        case "windows":
            result = WINDOWS_DEVICECLEANUPRULEPLATFORMTYPE
        case "windowsHolographic":
            result = WINDOWSHOLOGRAPHIC_DEVICECLEANUPRULEPLATFORMTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DEVICECLEANUPRULEPLATFORMTYPE
        default:
            return 0, errors.New("Unknown DeviceCleanupRulePlatformType value: " + v)
    }
    return &result, nil
}
func SerializeDeviceCleanupRulePlatformType(values []DeviceCleanupRulePlatformType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DeviceCleanupRulePlatformType) isMultiValue() bool {
    return false
}
