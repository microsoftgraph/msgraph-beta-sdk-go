package models
import (
    "errors"
)
// Specifies which platform(s) can be targeted for a given Android LOB application or Managed Android LOB application.
type AndroidTargetedPlatforms int

const (
    // Indicates the Android targeted platform is Android Device Administrator.
    ANDROIDDEVICEADMINISTRATOR_ANDROIDTARGETEDPLATFORMS AndroidTargetedPlatforms = iota
    // Indicates the Android targeted platform is Android Open Source Project.
    ANDROIDOPENSOURCEPROJECT_ANDROIDTARGETEDPLATFORMS
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_ANDROIDTARGETEDPLATFORMS
)

func (i AndroidTargetedPlatforms) String() string {
    return []string{"androidDeviceAdministrator", "androidOpenSourceProject", "unknownFutureValue"}[i]
}
func ParseAndroidTargetedPlatforms(v string) (any, error) {
    result := ANDROIDDEVICEADMINISTRATOR_ANDROIDTARGETEDPLATFORMS
    switch v {
        case "androidDeviceAdministrator":
            result = ANDROIDDEVICEADMINISTRATOR_ANDROIDTARGETEDPLATFORMS
        case "androidOpenSourceProject":
            result = ANDROIDOPENSOURCEPROJECT_ANDROIDTARGETEDPLATFORMS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ANDROIDTARGETEDPLATFORMS
        default:
            return 0, errors.New("Unknown AndroidTargetedPlatforms value: " + v)
    }
    return &result, nil
}
func SerializeAndroidTargetedPlatforms(values []AndroidTargetedPlatforms) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
