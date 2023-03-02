package models
import (
    "errors"
)
// The enrollment token type for an enrollment profile.
type AndroidDeviceOwnerEnrollmentTokenType int

const (
    // Default token type.
    DEFAULTESCAPED_ANDROIDDEVICEOWNERENROLLMENTTOKENTYPE AndroidDeviceOwnerEnrollmentTokenType = iota
    // Token type for Azure AD shared dedicated device enrollment. It applies to CorporateOwnedDedicatedDevice enrollment mode only.
    CORPORATEOWNEDDEDICATEDDEVICEWITHAZUREADSHAREDMODE_ANDROIDDEVICEOWNERENROLLMENTTOKENTYPE
)

func (i AndroidDeviceOwnerEnrollmentTokenType) String() string {
    return []string{"default", "corporateOwnedDedicatedDeviceWithAzureADSharedMode"}[i]
}
func ParseAndroidDeviceOwnerEnrollmentTokenType(v string) (any, error) {
    result := DEFAULTESCAPED_ANDROIDDEVICEOWNERENROLLMENTTOKENTYPE
    switch v {
        case "default":
            result = DEFAULTESCAPED_ANDROIDDEVICEOWNERENROLLMENTTOKENTYPE
        case "corporateOwnedDedicatedDeviceWithAzureADSharedMode":
            result = CORPORATEOWNEDDEDICATEDDEVICEWITHAZUREADSHAREDMODE_ANDROIDDEVICEOWNERENROLLMENTTOKENTYPE
        default:
            return 0, errors.New("Unknown AndroidDeviceOwnerEnrollmentTokenType value: " + v)
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
