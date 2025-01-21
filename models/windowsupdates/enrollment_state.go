package windowsupdates
type EnrollmentState int

const (
    NOTENROLLED_ENROLLMENTSTATE EnrollmentState = iota
    ENROLLED_ENROLLMENTSTATE
    ENROLLEDWITHPOLICY_ENROLLMENTSTATE
    ENROLLING_ENROLLMENTSTATE
    UNENROLLING_ENROLLMENTSTATE
    UNKNOWNFUTUREVALUE_ENROLLMENTSTATE
)

func (i EnrollmentState) String() string {
    return []string{"notEnrolled", "enrolled", "enrolledWithPolicy", "enrolling", "unenrolling", "unknownFutureValue"}[i]
}
func ParseEnrollmentState(v string) (any, error) {
    result := NOTENROLLED_ENROLLMENTSTATE
    switch v {
        case "notEnrolled":
            result = NOTENROLLED_ENROLLMENTSTATE
        case "enrolled":
            result = ENROLLED_ENROLLMENTSTATE
        case "enrolledWithPolicy":
            result = ENROLLEDWITHPOLICY_ENROLLMENTSTATE
        case "enrolling":
            result = ENROLLING_ENROLLMENTSTATE
        case "unenrolling":
            result = UNENROLLING_ENROLLMENTSTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ENROLLMENTSTATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeEnrollmentState(values []EnrollmentState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i EnrollmentState) isMultiValue() bool {
    return false
}
