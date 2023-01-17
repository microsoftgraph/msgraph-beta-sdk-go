package models
import (
    "errors"
)
// Android for Work device management targeting type for the account
type AndroidForWorkEnrollmentTarget int

const (
    NONE_ANDROIDFORWORKENROLLMENTTARGET AndroidForWorkEnrollmentTarget = iota
    ALL_ANDROIDFORWORKENROLLMENTTARGET
    TARGETED_ANDROIDFORWORKENROLLMENTTARGET
    TARGETEDASENROLLMENTRESTRICTIONS_ANDROIDFORWORKENROLLMENTTARGET
)

func (i AndroidForWorkEnrollmentTarget) String() string {
    return []string{"none", "all", "targeted", "targetedAsEnrollmentRestrictions"}[i]
}
func ParseAndroidForWorkEnrollmentTarget(v string) (any, error) {
    result := NONE_ANDROIDFORWORKENROLLMENTTARGET
    switch v {
        case "none":
            result = NONE_ANDROIDFORWORKENROLLMENTTARGET
        case "all":
            result = ALL_ANDROIDFORWORKENROLLMENTTARGET
        case "targeted":
            result = TARGETED_ANDROIDFORWORKENROLLMENTTARGET
        case "targetedAsEnrollmentRestrictions":
            result = TARGETEDASENROLLMENTRESTRICTIONS_ANDROIDFORWORKENROLLMENTTARGET
        default:
            return 0, errors.New("Unknown AndroidForWorkEnrollmentTarget value: " + v)
    }
    return &result, nil
}
func SerializeAndroidForWorkEnrollmentTarget(values []AndroidForWorkEnrollmentTarget) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
