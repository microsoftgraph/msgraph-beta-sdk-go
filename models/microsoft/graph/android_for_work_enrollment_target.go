package graph
import (
    "strings"
    "errors"
)
// 
type AndroidForWorkEnrollmentTarget int

const (
    NONE_ANDROIDFORWORKENROLLMENTTARGET AndroidForWorkEnrollmentTarget = iota
    ALL_ANDROIDFORWORKENROLLMENTTARGET
    TARGETED_ANDROIDFORWORKENROLLMENTTARGET
    TARGETEDASENROLLMENTRESTRICTIONS_ANDROIDFORWORKENROLLMENTTARGET
)

func (i AndroidForWorkEnrollmentTarget) String() string {
    return []string{"NONE", "ALL", "TARGETED", "TARGETEDASENROLLMENTRESTRICTIONS"}[i]
}
func ParseAndroidForWorkEnrollmentTarget(v string) (interface{}, error) {
    result := NONE_ANDROIDFORWORKENROLLMENTTARGET
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_ANDROIDFORWORKENROLLMENTTARGET
        case "ALL":
            result = ALL_ANDROIDFORWORKENROLLMENTTARGET
        case "TARGETED":
            result = TARGETED_ANDROIDFORWORKENROLLMENTTARGET
        case "TARGETEDASENROLLMENTRESTRICTIONS":
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
