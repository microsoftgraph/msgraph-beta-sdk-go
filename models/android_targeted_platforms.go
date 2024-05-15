package models
import (
    "math"
    "strings"
)
// Specifies which platform(s) can be targeted for a given Android LOB application or Managed Android LOB application.
type AndroidTargetedPlatforms int

const (
    // Indicates the Android targeted platform is Android Device Administrator.
    ANDROIDDEVICEADMINISTRATOR_ANDROIDTARGETEDPLATFORMS = 1
    // Indicates the Android targeted platform is Android Open Source Project.
    ANDROIDOPENSOURCEPROJECT_ANDROIDTARGETEDPLATFORMS = 2
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_ANDROIDTARGETEDPLATFORMS = 4
)

func (i AndroidTargetedPlatforms) String() string {
    var values []string
    options := []string{"androidDeviceAdministrator", "androidOpenSourceProject", "unknownFutureValue"}
    for p := 0; p < 3; p++ {
        mantis := AndroidTargetedPlatforms(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
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
                return nil, nil
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
