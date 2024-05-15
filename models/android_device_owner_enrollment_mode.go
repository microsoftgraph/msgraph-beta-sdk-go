package models
// The enrollment mode for an enrollment profile.
type AndroidDeviceOwnerEnrollmentMode int

const (
    CORPORATEOWNEDDEDICATEDDEVICE_ANDROIDDEVICEOWNERENROLLMENTMODE AndroidDeviceOwnerEnrollmentMode = iota
    CORPORATEOWNEDFULLYMANAGED_ANDROIDDEVICEOWNERENROLLMENTMODE
    CORPORATEOWNEDWORKPROFILE_ANDROIDDEVICEOWNERENROLLMENTMODE
    // Corporate owned, userless Android Open Source Project (AOSP) device, without Google Mobile Services.
    CORPORATEOWNEDAOSPUSERLESSDEVICE_ANDROIDDEVICEOWNERENROLLMENTMODE
    // Corporate owned, user-associated Android Open Source Project (AOSP) device, without Google Mobile Services.
    CORPORATEOWNEDAOSPUSERASSOCIATEDDEVICE_ANDROIDDEVICEOWNERENROLLMENTMODE
)

func (i AndroidDeviceOwnerEnrollmentMode) String() string {
    return []string{"corporateOwnedDedicatedDevice", "corporateOwnedFullyManaged", "corporateOwnedWorkProfile", "corporateOwnedAOSPUserlessDevice", "corporateOwnedAOSPUserAssociatedDevice"}[i]
}
func ParseAndroidDeviceOwnerEnrollmentMode(v string) (any, error) {
    result := CORPORATEOWNEDDEDICATEDDEVICE_ANDROIDDEVICEOWNERENROLLMENTMODE
    switch v {
        case "corporateOwnedDedicatedDevice":
            result = CORPORATEOWNEDDEDICATEDDEVICE_ANDROIDDEVICEOWNERENROLLMENTMODE
        case "corporateOwnedFullyManaged":
            result = CORPORATEOWNEDFULLYMANAGED_ANDROIDDEVICEOWNERENROLLMENTMODE
        case "corporateOwnedWorkProfile":
            result = CORPORATEOWNEDWORKPROFILE_ANDROIDDEVICEOWNERENROLLMENTMODE
        case "corporateOwnedAOSPUserlessDevice":
            result = CORPORATEOWNEDAOSPUSERLESSDEVICE_ANDROIDDEVICEOWNERENROLLMENTMODE
        case "corporateOwnedAOSPUserAssociatedDevice":
            result = CORPORATEOWNEDAOSPUSERASSOCIATEDDEVICE_ANDROIDDEVICEOWNERENROLLMENTMODE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAndroidDeviceOwnerEnrollmentMode(values []AndroidDeviceOwnerEnrollmentMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AndroidDeviceOwnerEnrollmentMode) isMultiValue() bool {
    return false
}
