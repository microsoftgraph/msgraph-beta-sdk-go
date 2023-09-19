package models
import (
    "errors"
    "strings"
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
    var values []string
    for p := AndroidTargetedPlatforms(1); p <= UNKNOWNFUTUREVALUE_ANDROIDTARGETEDPLATFORMS; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"androidDeviceAdministrator", "androidOpenSourceProject", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseAndroidTargetedPlatforms(v string) (any, error) {
    var result AndroidTargetedPlatforms
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "androidDeviceAdministrator":
                result |= ANDROIDDEVICEADMINISTRATOR_ANDROIDTARGETEDPLATFORMS
            case "androidOpenSourceProject":
                result |= ANDROIDOPENSOURCEPROJECT_ANDROIDTARGETEDPLATFORMS
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_ANDROIDTARGETEDPLATFORMS
            default:
                return 0, errors.New("Unknown AndroidTargetedPlatforms value: " + v)
        }
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
func (i AndroidTargetedPlatforms) isMultiValue() bool {
    return true
}
