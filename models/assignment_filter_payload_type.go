package models
// Represents the payload type AssignmentFilter is being assigned to.
type AssignmentFilterPayloadType int

const (
    // NotSet
    NOTSET_ASSIGNMENTFILTERPAYLOADTYPE AssignmentFilterPayloadType = iota
    // EnrollmentRestrictions
    ENROLLMENTRESTRICTIONS_ASSIGNMENTFILTERPAYLOADTYPE
)

func (i AssignmentFilterPayloadType) String() string {
    return []string{"notSet", "enrollmentRestrictions"}[i]
}
func ParseAssignmentFilterPayloadType(v string) (any, error) {
    result := NOTSET_ASSIGNMENTFILTERPAYLOADTYPE
    switch v {
        case "notSet":
            result = NOTSET_ASSIGNMENTFILTERPAYLOADTYPE
        case "enrollmentRestrictions":
            result = ENROLLMENTRESTRICTIONS_ASSIGNMENTFILTERPAYLOADTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAssignmentFilterPayloadType(values []AssignmentFilterPayloadType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AssignmentFilterPayloadType) isMultiValue() bool {
    return false
}
