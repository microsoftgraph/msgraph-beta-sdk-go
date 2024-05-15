package models
// An enum representing possible values for cross profile data sharing.
type AndroidDeviceOwnerCrossProfileDataSharing int

const (
    // Not configured; this value defaults to CROSS_PROFILE_DATA_SHARING_UNSPECIFIED.
    NOTCONFIGURED_ANDROIDDEVICEOWNERCROSSPROFILEDATASHARING AndroidDeviceOwnerCrossProfileDataSharing = iota
    // Data cannot be shared from both the personal profile to work profile and the work profile to the personal profile.
    CROSSPROFILEDATASHARINGBLOCKED_ANDROIDDEVICEOWNERCROSSPROFILEDATASHARING
    // Prevents users from sharing data from the work profile to apps in the personal profile. Personal data can be shared with work apps.
    DATASHARINGFROMWORKTOPERSONALBLOCKED_ANDROIDDEVICEOWNERCROSSPROFILEDATASHARING
    // Data from either profile can be shared with the other profile.
    CROSSPROFILEDATASHARINGALLOWED_ANDROIDDEVICEOWNERCROSSPROFILEDATASHARING
    // Unknown future value (reserved, not used right now)
    UNKOWNFUTUREVALUE_ANDROIDDEVICEOWNERCROSSPROFILEDATASHARING
)

func (i AndroidDeviceOwnerCrossProfileDataSharing) String() string {
    return []string{"notConfigured", "crossProfileDataSharingBlocked", "dataSharingFromWorkToPersonalBlocked", "crossProfileDataSharingAllowed", "unkownFutureValue"}[i]
}
func ParseAndroidDeviceOwnerCrossProfileDataSharing(v string) (any, error) {
    result := NOTCONFIGURED_ANDROIDDEVICEOWNERCROSSPROFILEDATASHARING
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_ANDROIDDEVICEOWNERCROSSPROFILEDATASHARING
        case "crossProfileDataSharingBlocked":
            result = CROSSPROFILEDATASHARINGBLOCKED_ANDROIDDEVICEOWNERCROSSPROFILEDATASHARING
        case "dataSharingFromWorkToPersonalBlocked":
            result = DATASHARINGFROMWORKTOPERSONALBLOCKED_ANDROIDDEVICEOWNERCROSSPROFILEDATASHARING
        case "crossProfileDataSharingAllowed":
            result = CROSSPROFILEDATASHARINGALLOWED_ANDROIDDEVICEOWNERCROSSPROFILEDATASHARING
        case "unkownFutureValue":
            result = UNKOWNFUTUREVALUE_ANDROIDDEVICEOWNERCROSSPROFILEDATASHARING
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAndroidDeviceOwnerCrossProfileDataSharing(values []AndroidDeviceOwnerCrossProfileDataSharing) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AndroidDeviceOwnerCrossProfileDataSharing) isMultiValue() bool {
    return false
}
