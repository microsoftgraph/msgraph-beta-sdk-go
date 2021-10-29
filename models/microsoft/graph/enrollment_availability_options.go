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
    switch strings.ToUpper(v) {
        case "AVAILABLEWITHPROMPTS":
            return AVAILABLEWITHPROMPTS_ENROLLMENTAVAILABILITYOPTIONS, nil
        case "AVAILABLEWITHOUTPROMPTS":
            return AVAILABLEWITHOUTPROMPTS_ENROLLMENTAVAILABILITYOPTIONS, nil
        case "UNAVAILABLE":
            return UNAVAILABLE_ENROLLMENTAVAILABILITYOPTIONS, nil
    }
    return 0, errors.New("Unknown EnrollmentAvailabilityOptions value: " + v)
}
func SerializeEnrollmentAvailabilityOptions(values []EnrollmentAvailabilityOptions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
