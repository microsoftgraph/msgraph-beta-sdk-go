package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_ANDROIDFORWORKENROLLMENTTARGET, nil
        case "ALL":
            return ALL_ANDROIDFORWORKENROLLMENTTARGET, nil
        case "TARGETED":
            return TARGETED_ANDROIDFORWORKENROLLMENTTARGET, nil
        case "TARGETEDASENROLLMENTRESTRICTIONS":
            return TARGETEDASENROLLMENTRESTRICTIONS_ANDROIDFORWORKENROLLMENTTARGET, nil
    }
    return 0, errors.New("Unknown AndroidForWorkEnrollmentTarget value: " + v)
}
func SerializeAndroidForWorkEnrollmentTarget(values []AndroidForWorkEnrollmentTarget) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
