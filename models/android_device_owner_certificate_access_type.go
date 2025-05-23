// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models
// An enum representing possible values for cross profile data sharing.
type AndroidDeviceOwnerCertificateAccessType int

const (
    // Require user approval for all apps
    USERAPPROVAL_ANDROIDDEVICEOWNERCERTIFICATEACCESSTYPE AndroidDeviceOwnerCertificateAccessType = iota
    // Pre-grant certificate access for specific apps (require user approval for other apps).
    SPECIFICAPPS_ANDROIDDEVICEOWNERCERTIFICATEACCESSTYPE
    // Evolvable enumeration sentinel value. Do not use.
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
            return nil, nil
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
func (i AndroidDeviceOwnerCertificateAccessType) isMultiValue() bool {
    return false
}
