package models
import (
    "errors"
)
// An enum representing possible values for cross profile data sharing.
type AndroidDeviceOwnerCertificateAccessType int

const (
    // Require user approval for all apps
    USERAPPROVAL_ANDROIDDEVICEOWNERCERTIFICATEACCESSTYPE AndroidDeviceOwnerCertificateAccessType = iota
    // Pre-grant certificate access for specific apps (require user approval for other apps).
    SPECIFICAPPS_ANDROIDDEVICEOWNERCERTIFICATEACCESSTYPE
    // Unknown future value for evolvable enum patterns.
    UNKNOWNFUTUREVALUE_ANDROIDDEVICEOWNERCERTIFICATEACCESSTYPE
)

func (i AndroidDeviceOwnerCertificateAccessType) String() string {
    return []string{"userApproval", "specificApps", "unknownFutureValue"}[i]
}
func ParseAndroidDeviceOwnerCertificateAccessType(v string) (any, error) {
    result := USERAPPROVAL_ANDROIDDEVICEOWNERCERTIFICATEACCESSTYPE
    switch v {
        case "userApproval":
            result = USERAPPROVAL_ANDROIDDEVICEOWNERCERTIFICATEACCESSTYPE
        case "specificApps":
            result = SPECIFICAPPS_ANDROIDDEVICEOWNERCERTIFICATEACCESSTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ANDROIDDEVICEOWNERCERTIFICATEACCESSTYPE
        default:
            return 0, errors.New("Unknown AndroidDeviceOwnerCertificateAccessType value: " + v)
    }
    return &result, nil
}
func SerializeAndroidDeviceOwnerCertificateAccessType(values []AndroidDeviceOwnerCertificateAccessType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
