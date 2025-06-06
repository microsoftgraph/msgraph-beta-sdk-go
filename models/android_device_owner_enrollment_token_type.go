// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models
// The enrollment token type for an enrollment profile.
type AndroidDeviceOwnerEnrollmentTokenType int

const (
    // Default token type.
    DEFAULT_ANDROIDDEVICEOWNERENROLLMENTTOKENTYPE AndroidDeviceOwnerEnrollmentTokenType = iota
    // Token type for Azure AD shared dedicated device enrollment. It applies to CorporateOwnedDedicatedDevice enrollment mode only.
    CORPORATEOWNEDDEDICATEDDEVICEWITHAZUREADSHAREDMODE_ANDROIDDEVICEOWNERENROLLMENTTOKENTYPE
    // Token type for Android Device Staging enrollment type. It applies to CorporateOwnedFullyManaged and CorporateOwnedWorkProfile only.
    DEVICESTAGING_ANDROIDDEVICEOWNERENROLLMENTTOKENTYPE
)

func (i AndroidDeviceOwnerEnrollmentTokenType) String() string {
    return []string{"default", "corporateOwnedDedicatedDeviceWithAzureADSharedMode", "deviceStaging"}[i]
}
func ParseAndroidDeviceOwnerEnrollmentTokenType(v string) (any, error) {
    result := DEFAULT_ANDROIDDEVICEOWNERENROLLMENTTOKENTYPE
    switch v {
        case "default":
            result = DEFAULT_ANDROIDDEVICEOWNERENROLLMENTTOKENTYPE
        case "corporateOwnedDedicatedDeviceWithAzureADSharedMode":
            result = CORPORATEOWNEDDEDICATEDDEVICEWITHAZUREADSHAREDMODE_ANDROIDDEVICEOWNERENROLLMENTTOKENTYPE
        case "deviceStaging":
            result = DEVICESTAGING_ANDROIDDEVICEOWNERENROLLMENTTOKENTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAndroidDeviceOwnerEnrollmentTokenType(values []AndroidDeviceOwnerEnrollmentTokenType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AndroidDeviceOwnerEnrollmentTokenType) isMultiValue() bool {
    return false
}
