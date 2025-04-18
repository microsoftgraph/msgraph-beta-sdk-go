// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models
// Indicates the device licensing status after Windows device based subscription has been enabled.
type DeviceLicensingStatus int

const (
    // Default. Set to unknown when status cannot be determined.
    UNKNOWN_DEVICELICENSINGSTATUS DeviceLicensingStatus = iota
    // This status is set when the license refresh is started.
    LICENSEREFRESHSTARTED_DEVICELICENSINGSTATUS
    // This status is set when the license refresh is pending.
    LICENSEREFRESHPENDING_DEVICELICENSINGSTATUS
    // This status is set when the device is not joined to Azure Active Directory.
    DEVICEISNOTAZUREACTIVEDIRECTORYJOINED_DEVICELICENSINGSTATUS
    // This status is set when the Microsoft device identity is being verified.
    VERIFYINGMICROSOFTDEVICEIDENTITY_DEVICELICENSINGSTATUS
    // This status is set when the Microsoft device identity verification fails.
    DEVICEIDENTITYVERIFICATIONFAILED_DEVICELICENSINGSTATUS
    // This status is set when the Microsoft account identity is being verified.
    VERIFYINGMICROSOFTACCOUNTIDENTITY_DEVICELICENSINGSTATUS
    // This status is set when the Microsoft account identity verification fails.
    MICROSOFTACCOUNTVERIFICATIONFAILED_DEVICELICENSINGSTATUS
    // This status is set when the device license is being acquired.
    ACQUIRINGDEVICELICENSE_DEVICELICENSINGSTATUS
    // This status is set when the device license is being refreshed.
    REFRESHINGDEVICELICENSE_DEVICELICENSINGSTATUS
    // This status is set when the device license refresh succeeds.
    DEVICELICENSEREFRESHSUCCEED_DEVICELICENSINGSTATUS
    // This status is set when the device license refresh fails.
    DEVICELICENSEREFRESHFAILED_DEVICELICENSINGSTATUS
    // This status is set when the device license is being removed.
    REMOVINGDEVICELICENSE_DEVICELICENSINGSTATUS
    // This status is set when the device license removing succeeds.
    DEVICELICENSEREMOVESUCCEED_DEVICELICENSINGSTATUS
    // This status is set when the device license removing fails.
    DEVICELICENSEREMOVEFAILED_DEVICELICENSINGSTATUS
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_DEVICELICENSINGSTATUS
)

func (i DeviceLicensingStatus) String() string {
    return []string{"unknown", "licenseRefreshStarted", "licenseRefreshPending", "deviceIsNotAzureActiveDirectoryJoined", "verifyingMicrosoftDeviceIdentity", "deviceIdentityVerificationFailed", "verifyingMicrosoftAccountIdentity", "microsoftAccountVerificationFailed", "acquiringDeviceLicense", "refreshingDeviceLicense", "deviceLicenseRefreshSucceed", "deviceLicenseRefreshFailed", "removingDeviceLicense", "deviceLicenseRemoveSucceed", "deviceLicenseRemoveFailed", "unknownFutureValue"}[i]
}
func ParseDeviceLicensingStatus(v string) (any, error) {
    result := UNKNOWN_DEVICELICENSINGSTATUS
    switch v {
        case "unknown":
            result = UNKNOWN_DEVICELICENSINGSTATUS
        case "licenseRefreshStarted":
            result = LICENSEREFRESHSTARTED_DEVICELICENSINGSTATUS
        case "licenseRefreshPending":
            result = LICENSEREFRESHPENDING_DEVICELICENSINGSTATUS
        case "deviceIsNotAzureActiveDirectoryJoined":
            result = DEVICEISNOTAZUREACTIVEDIRECTORYJOINED_DEVICELICENSINGSTATUS
        case "verifyingMicrosoftDeviceIdentity":
            result = VERIFYINGMICROSOFTDEVICEIDENTITY_DEVICELICENSINGSTATUS
        case "deviceIdentityVerificationFailed":
            result = DEVICEIDENTITYVERIFICATIONFAILED_DEVICELICENSINGSTATUS
        case "verifyingMicrosoftAccountIdentity":
            result = VERIFYINGMICROSOFTACCOUNTIDENTITY_DEVICELICENSINGSTATUS
        case "microsoftAccountVerificationFailed":
            result = MICROSOFTACCOUNTVERIFICATIONFAILED_DEVICELICENSINGSTATUS
        case "acquiringDeviceLicense":
            result = ACQUIRINGDEVICELICENSE_DEVICELICENSINGSTATUS
        case "refreshingDeviceLicense":
            result = REFRESHINGDEVICELICENSE_DEVICELICENSINGSTATUS
        case "deviceLicenseRefreshSucceed":
            result = DEVICELICENSEREFRESHSUCCEED_DEVICELICENSINGSTATUS
        case "deviceLicenseRefreshFailed":
            result = DEVICELICENSEREFRESHFAILED_DEVICELICENSINGSTATUS
        case "removingDeviceLicense":
            result = REMOVINGDEVICELICENSE_DEVICELICENSINGSTATUS
        case "deviceLicenseRemoveSucceed":
            result = DEVICELICENSEREMOVESUCCEED_DEVICELICENSINGSTATUS
        case "deviceLicenseRemoveFailed":
            result = DEVICELICENSEREMOVEFAILED_DEVICELICENSINGSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DEVICELICENSINGSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDeviceLicensingStatus(values []DeviceLicensingStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DeviceLicensingStatus) isMultiValue() bool {
    return false
}
