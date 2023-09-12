package models
import (
    "errors"
)
// An enum representing possible values for delegated app scope.
type AndroidDeviceOwnerDelegatedAppScopeType int

const (
    // Unspecified; this value defaults to DELEGATED_SCOPE_UNSPECIFIED.
    UNSPECIFIED_ANDROIDDEVICEOWNERDELEGATEDAPPSCOPETYPE AndroidDeviceOwnerDelegatedAppScopeType = iota
    // Specifies that the admin has given app permission to install and manage certificates on device.
    CERTIFICATEINSTALL_ANDROIDDEVICEOWNERDELEGATEDAPPSCOPETYPE
    // Specifies that the admin has given app permission to capture network activity logs on device. More info on Network activity logs: https://developer.android.com/work/dpc/logging 
    CAPTURENETWORKACTIVITYLOG_ANDROIDDEVICEOWNERDELEGATEDAPPSCOPETYPE
    // Specified that the admin has given permission to capture security logs on device. More info on Security logs: https://developer.android.com/work/dpc/security#log_enterprise_device_activity
    CAPTURESECURITYLOG_ANDROIDDEVICEOWNERDELEGATEDAPPSCOPETYPE
    // Unknown future value (reserved, not used right now)
    UNKNOWNFUTUREVALUE_ANDROIDDEVICEOWNERDELEGATEDAPPSCOPETYPE
)

func (i AndroidDeviceOwnerDelegatedAppScopeType) String() string {
    return []string{"unspecified", "certificateInstall", "captureNetworkActivityLog", "captureSecurityLog", "unknownFutureValue"}[i]
}
func ParseAndroidDeviceOwnerDelegatedAppScopeType(v string) (any, error) {
    result := UNSPECIFIED_ANDROIDDEVICEOWNERDELEGATEDAPPSCOPETYPE
    switch v {
        case "unspecified":
            result = UNSPECIFIED_ANDROIDDEVICEOWNERDELEGATEDAPPSCOPETYPE
        case "certificateInstall":
            result = CERTIFICATEINSTALL_ANDROIDDEVICEOWNERDELEGATEDAPPSCOPETYPE
        case "captureNetworkActivityLog":
            result = CAPTURENETWORKACTIVITYLOG_ANDROIDDEVICEOWNERDELEGATEDAPPSCOPETYPE
        case "captureSecurityLog":
            result = CAPTURESECURITYLOG_ANDROIDDEVICEOWNERDELEGATEDAPPSCOPETYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ANDROIDDEVICEOWNERDELEGATEDAPPSCOPETYPE
        default:
            return 0, errors.New("Unknown AndroidDeviceOwnerDelegatedAppScopeType value: " + v)
    }
    return &result, nil
}
func SerializeAndroidDeviceOwnerDelegatedAppScopeType(values []AndroidDeviceOwnerDelegatedAppScopeType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AndroidDeviceOwnerDelegatedAppScopeType) isMultiValue() bool {
    return false
}
