package graph
import (
    "strings"
    "errors"
)
// 
type AppleUserInitiatedEnrollmentType int

const (
    UNKNOWN_APPLEUSERINITIATEDENROLLMENTTYPE AppleUserInitiatedEnrollmentType = iota
    DEVICE_APPLEUSERINITIATEDENROLLMENTTYPE
    USER_APPLEUSERINITIATEDENROLLMENTTYPE
)

func (i AppleUserInitiatedEnrollmentType) String() string {
    return []string{"UNKNOWN", "DEVICE", "USER"}[i]
}
func ParseAppleUserInitiatedEnrollmentType(v string) (interface{}, error) {
    result := UNKNOWN_APPLEUSERINITIATEDENROLLMENTTYPE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_APPLEUSERINITIATEDENROLLMENTTYPE
        case "DEVICE":
            result = DEVICE_APPLEUSERINITIATEDENROLLMENTTYPE
        case "USER":
            result = USER_APPLEUSERINITIATEDENROLLMENTTYPE
        default:
            return 0, errors.New("Unknown AppleUserInitiatedEnrollmentType value: " + v)
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
