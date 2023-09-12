package models
import (
    "errors"
)
// Android profile applicability
type AndroidProfileApplicability int

const (
    DEFAULTESCAPED_ANDROIDPROFILEAPPLICABILITY AndroidProfileApplicability = iota
    ANDROIDWORKPROFILE_ANDROIDPROFILEAPPLICABILITY
    ANDROIDDEVICEOWNER_ANDROIDPROFILEAPPLICABILITY
)

func (i AndroidProfileApplicability) String() string {
    return []string{"default", "androidWorkProfile", "androidDeviceOwner"}[i]
}
func ParseAndroidProfileApplicability(v string) (any, error) {
    result := DEFAULTESCAPED_ANDROIDPROFILEAPPLICABILITY
    switch v {
        case "default":
            result = DEFAULTESCAPED_ANDROIDPROFILEAPPLICABILITY
        case "androidWorkProfile":
            result = ANDROIDWORKPROFILE_ANDROIDPROFILEAPPLICABILITY
        case "androidDeviceOwner":
            result = ANDROIDDEVICEOWNER_ANDROIDPROFILEAPPLICABILITY
        default:
            return 0, errors.New("Unknown AndroidProfileApplicability value: " + v)
    }
    return &result, nil
}
func SerializeAndroidProfileApplicability(values []AndroidProfileApplicability) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AndroidProfileApplicability) isMultiValue() bool {
    return false
}
