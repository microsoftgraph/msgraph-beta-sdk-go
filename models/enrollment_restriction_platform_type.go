package models
import (
    "errors"
)
// Provides operations to manage the collection of agreementAcceptance entities.
type EnrollmentRestrictionPlatformType int

const (
    // Applies to all platforms
    ALLPLATFORMS_ENROLLMENTRESTRICTIONPLATFORMTYPE EnrollmentRestrictionPlatformType = iota
    // iOS/iPadOS devices
    IOS_ENROLLMENTRESTRICTIONPLATFORMTYPE
    // Windows devices
    WINDOWS_ENROLLMENTRESTRICTIONPLATFORMTYPE
    // Windows Phone devices
    WINDOWSPHONE_ENROLLMENTRESTRICTIONPLATFORMTYPE
    // Android devices
    ANDROID_ENROLLMENTRESTRICTIONPLATFORMTYPE
    // Android for Work devices
    ANDROIDFORWORK_ENROLLMENTRESTRICTIONPLATFORMTYPE
    // macOS devices
    MAC_ENROLLMENTRESTRICTIONPLATFORMTYPE
)

func (i EnrollmentRestrictionPlatformType) String() string {
    return []string{"allPlatforms", "ios", "windows", "windowsPhone", "android", "androidForWork", "mac"}[i]
}
func ParseEnrollmentRestrictionPlatformType(v string) (interface{}, error) {
    result := ALLPLATFORMS_ENROLLMENTRESTRICTIONPLATFORMTYPE
    switch v {
        case "allPlatforms":
            result = ALLPLATFORMS_ENROLLMENTRESTRICTIONPLATFORMTYPE
        case "ios":
            result = IOS_ENROLLMENTRESTRICTIONPLATFORMTYPE
        case "windows":
            result = WINDOWS_ENROLLMENTRESTRICTIONPLATFORMTYPE
        case "windowsPhone":
            result = WINDOWSPHONE_ENROLLMENTRESTRICTIONPLATFORMTYPE
        case "android":
            result = ANDROID_ENROLLMENTRESTRICTIONPLATFORMTYPE
        case "androidForWork":
            result = ANDROIDFORWORK_ENROLLMENTRESTRICTIONPLATFORMTYPE
        case "mac":
            result = MAC_ENROLLMENTRESTRICTIONPLATFORMTYPE
        default:
            return 0, errors.New("Unknown EnrollmentRestrictionPlatformType value: " + v)
    }
    return &result, nil
}
func SerializeEnrollmentRestrictionPlatformType(values []EnrollmentRestrictionPlatformType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
