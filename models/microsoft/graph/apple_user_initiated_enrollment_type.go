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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_APPLEUSERINITIATEDENROLLMENTTYPE, nil
        case "DEVICE":
            return DEVICE_APPLEUSERINITIATEDENROLLMENTTYPE, nil
        case "USER":
            return USER_APPLEUSERINITIATEDENROLLMENTTYPE, nil
    }
    return 0, errors.New("Unknown AppleUserInitiatedEnrollmentType value: " + v)
}
func SerializeAppleUserInitiatedEnrollmentType(values []AppleUserInitiatedEnrollmentType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
