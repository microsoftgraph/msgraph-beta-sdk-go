package models
type AppleUserInitiatedEnrollmentType int

const (
    // Default value in case enum parsing fails
    UNKNOWN_APPLEUSERINITIATEDENROLLMENTTYPE AppleUserInitiatedEnrollmentType = iota
    // Device enrollment via the iOS Company Portal. The default user-initiated enrollment type, which does not segregate corporate and personal data. Supported on all Intune-supported iOS/iPadOS versions.
    DEVICE_APPLEUSERINITIATEDENROLLMENTTYPE
    // Profile-driven user enrollment via the iOS Company Portal. An enrollment type that segregates corportate and personal data. Supported on devices running iOS/iPadOS 13 and higher.
    USER_APPLEUSERINITIATEDENROLLMENTTYPE
    // Account-driven user enrollment. Users will enroll from the iOS Settings app without using the iOS Company Portal. This enrollment type segregates corporate and personal data. Supported on devices running iOS/iPadOS 15 and higher.
    ACCOUNTDRIVENUSERENROLLMENT_APPLEUSERINITIATEDENROLLMENTTYPE
    // Device enrollment via the web. Users will enroll without using the iOS Company Portal. This enrollment type does not segregate corporate and personal data. Supported on all Intune-supported iOS/iPadOS versions.
    WEBDEVICEENROLLMENT_APPLEUSERINITIATEDENROLLMENTTYPE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_APPLEUSERINITIATEDENROLLMENTTYPE
)

func (i AppleUserInitiatedEnrollmentType) String() string {
    return []string{"unknown", "device", "user", "accountDrivenUserEnrollment", "webDeviceEnrollment", "unknownFutureValue"}[i]
}
func ParseAppleUserInitiatedEnrollmentType(v string) (any, error) {
    result := UNKNOWN_APPLEUSERINITIATEDENROLLMENTTYPE
    switch v {
        case "unknown":
            result = UNKNOWN_APPLEUSERINITIATEDENROLLMENTTYPE
        case "device":
            result = DEVICE_APPLEUSERINITIATEDENROLLMENTTYPE
        case "user":
            result = USER_APPLEUSERINITIATEDENROLLMENTTYPE
        case "accountDrivenUserEnrollment":
            result = ACCOUNTDRIVENUSERENROLLMENT_APPLEUSERINITIATEDENROLLMENTTYPE
        case "webDeviceEnrollment":
            result = WEBDEVICEENROLLMENT_APPLEUSERINITIATEDENROLLMENTTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_APPLEUSERINITIATEDENROLLMENTTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAppleUserInitiatedEnrollmentType(values []AppleUserInitiatedEnrollmentType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppleUserInitiatedEnrollmentType) isMultiValue() bool {
    return false
}
