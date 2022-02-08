package graph
import (
    "strings"
    "errors"
)
// 
type EnrollmentAvailabilityOptions int

const (
    AVAILABLEWITHPROMPTS_ENROLLMENTAVAILABILITYOPTIONS EnrollmentAvailabilityOptions = iota
    AVAILABLEWITHOUTPROMPTS_ENROLLMENTAVAILABILITYOPTIONS
    UNAVAILABLE_ENROLLMENTAVAILABILITYOPTIONS
)

func (i EnrollmentAvailabilityOptions) String() string {
    return []string{"AVAILABLEWITHPROMPTS", "AVAILABLEWITHOUTPROMPTS", "UNAVAILABLE"}[i]
}
func ParseEnrollmentAvailabilityOptions(v string) (interface{}, error) {
    result := AVAILABLEWITHPROMPTS_ENROLLMENTAVAILABILITYOPTIONS
    switch strings.ToUpper(v) {
        case "AVAILABLEWITHPROMPTS":
            result = AVAILABLEWITHPROMPTS_ENROLLMENTAVAILABILITYOPTIONS
        case "AVAILABLEWITHOUTPROMPTS":
            result = AVAILABLEWITHOUTPROMPTS_ENROLLMENTAVAILABILITYOPTIONS
        case "UNAVAILABLE":
            result = UNAVAILABLE_ENROLLMENTAVAILABILITYOPTIONS
        default:
            return 0, errors.New("Unknown EnrollmentAvailabilityOptions value: " + v)
    }
    return &result, nil
}
func SerializeEnrollmentAvailabilityOptions(values []EnrollmentAvailabilityOptions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
