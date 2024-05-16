package models
// Options available for enrollment flow customization
type EnrollmentAvailabilityOptions int

const (
    // Device enrollment flow is shown to the end user with guided enrollment prompts
    AVAILABLEWITHPROMPTS_ENROLLMENTAVAILABILITYOPTIONS EnrollmentAvailabilityOptions = iota
    // Device enrollment flow is available to the end user without guided enrollment prompts
    AVAILABLEWITHOUTPROMPTS_ENROLLMENTAVAILABILITYOPTIONS
    // Device enrollment flow is unavailable to the enduser
    UNAVAILABLE_ENROLLMENTAVAILABILITYOPTIONS
)

func (i EnrollmentAvailabilityOptions) String() string {
    return []string{"availableWithPrompts", "availableWithoutPrompts", "unavailable"}[i]
}
func ParseEnrollmentAvailabilityOptions(v string) (any, error) {
    result := AVAILABLEWITHPROMPTS_ENROLLMENTAVAILABILITYOPTIONS
    switch v {
        case "availableWithPrompts":
            result = AVAILABLEWITHPROMPTS_ENROLLMENTAVAILABILITYOPTIONS
        case "availableWithoutPrompts":
            result = AVAILABLEWITHOUTPROMPTS_ENROLLMENTAVAILABILITYOPTIONS
        case "unavailable":
            result = UNAVAILABLE_ENROLLMENTAVAILABILITYOPTIONS
        default:
            return nil, nil
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
func (i EnrollmentAvailabilityOptions) isMultiValue() bool {
    return false
}
