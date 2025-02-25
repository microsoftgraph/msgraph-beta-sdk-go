package models
import (
    "math"
    "strings"
)
// Management levels for apps
type AppManagementLevel int

const (
    // Unspecified
    UNSPECIFIED_APPMANAGEMENTLEVEL = 1
    // Unmanaged
    UNMANAGED_APPMANAGEMENTLEVEL = 2
    // MDM
    MDM_APPMANAGEMENTLEVEL = 4
    // Android Enterprise
    ANDROIDENTERPRISE_APPMANAGEMENTLEVEL = 8
    // Android Enterprise dedicated devices with Azure AD Shared mode
    ANDROIDENTERPRISEDEDICATEDDEVICESWITHAZUREADSHAREDMODE_APPMANAGEMENTLEVEL = 16
    // Android Open Source Project (AOSP) devices
    ANDROIDOPENSOURCEPROJECTUSERASSOCIATED_APPMANAGEMENTLEVEL = 32
    // Android Open Source Project (AOSP) userless devices
    ANDROIDOPENSOURCEPROJECTUSERLESS_APPMANAGEMENTLEVEL = 64
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_APPMANAGEMENTLEVEL = 128
)

func (i AppManagementLevel) String() string {
    var values []string
    options := []string{"unspecified", "unmanaged", "mdm", "androidEnterprise", "androidEnterpriseDedicatedDevicesWithAzureAdSharedMode", "androidOpenSourceProjectUserAssociated", "androidOpenSourceProjectUserless", "unknownFutureValue"}
    for p := 0; p < 8; p++ {
        mantis := AppManagementLevel(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseAppManagementLevel(v string) (any, error) {
    var result AppManagementLevel
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "unspecified":
                result |= UNSPECIFIED_APPMANAGEMENTLEVEL
            case "unmanaged":
                result |= UNMANAGED_APPMANAGEMENTLEVEL
            case "mdm":
                result |= MDM_APPMANAGEMENTLEVEL
            case "androidEnterprise":
                result |= ANDROIDENTERPRISE_APPMANAGEMENTLEVEL
            case "androidEnterpriseDedicatedDevicesWithAzureAdSharedMode":
                result |= ANDROIDENTERPRISEDEDICATEDDEVICESWITHAZUREADSHAREDMODE_APPMANAGEMENTLEVEL
            case "androidOpenSourceProjectUserAssociated":
                result |= ANDROIDOPENSOURCEPROJECTUSERASSOCIATED_APPMANAGEMENTLEVEL
            case "androidOpenSourceProjectUserless":
                result |= ANDROIDOPENSOURCEPROJECTUSERLESS_APPMANAGEMENTLEVEL
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_APPMANAGEMENTLEVEL
            default:
                return nil, nil
        }
    }
    return &result, nil
}
func SerializeAppManagementLevel(values []AppManagementLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppManagementLevel) isMultiValue() bool {
    return true
}
