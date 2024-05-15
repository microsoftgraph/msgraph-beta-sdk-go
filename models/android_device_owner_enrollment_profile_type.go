package models
// Android Device Owner Enrollment Profile types.
type AndroidDeviceOwnerEnrollmentProfileType int

const (
    // Not configured; this value is ignored.
    NOTCONFIGURED_ANDROIDDEVICEOWNERENROLLMENTPROFILETYPE AndroidDeviceOwnerEnrollmentProfileType = iota
    // Dedicated device.
    DEDICATEDDEVICE_ANDROIDDEVICEOWNERENROLLMENTPROFILETYPE
    // Fully managed.
    FULLYMANAGED_ANDROIDDEVICEOWNERENROLLMENTPROFILETYPE
)

func (i AndroidDeviceOwnerEnrollmentProfileType) String() string {
    return []string{"notConfigured", "dedicatedDevice", "fullyManaged"}[i]
}
func ParseAndroidDeviceOwnerEnrollmentProfileType(v string) (any, error) {
    result := NOTCONFIGURED_ANDROIDDEVICEOWNERENROLLMENTPROFILETYPE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_ANDROIDDEVICEOWNERENROLLMENTPROFILETYPE
        case "dedicatedDevice":
            result = DEDICATEDDEVICE_ANDROIDDEVICEOWNERENROLLMENTPROFILETYPE
        case "fullyManaged":
            result = FULLYMANAGED_ANDROIDDEVICEOWNERENROLLMENTPROFILETYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAndroidDeviceOwnerEnrollmentProfileType(values []AndroidDeviceOwnerEnrollmentProfileType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AndroidDeviceOwnerEnrollmentProfileType) isMultiValue() bool {
    return false
}
