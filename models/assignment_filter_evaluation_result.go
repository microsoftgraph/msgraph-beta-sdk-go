package models
// Supported evaluation results for filter.
type AssignmentFilterEvaluationResult int

const (
    // Unknown.
    UNKNOWN_ASSIGNMENTFILTEREVALUATIONRESULT AssignmentFilterEvaluationResult = iota
    // Match.
    MATCH_ASSIGNMENTFILTEREVALUATIONRESULT
    // NotMatch.
    NOTMATCH_ASSIGNMENTFILTEREVALUATIONRESULT
    // Inconclusive.
    INCONCLUSIVE_ASSIGNMENTFILTEREVALUATIONRESULT
    // Failure.
    FAILURE_ASSIGNMENTFILTEREVALUATIONRESULT
    // NotEvaluated.
    NOTEVALUATED_ASSIGNMENTFILTEREVALUATIONRESULT
)

func (i AssignmentFilterEvaluationResult) String() string {
    return []string{"unknown", "match", "notMatch", "inconclusive", "failure", "notEvaluated"}[i]
}
func ParseAssignmentFilterEvaluationResult(v string) (any, error) {
    result := UNKNOWN_ASSIGNMENTFILTEREVALUATIONRESULT
    switch v {
        case "unknown":
            result = UNKNOWN_ASSIGNMENTFILTEREVALUATIONRESULT
        case "match":
            result = MATCH_ASSIGNMENTFILTEREVALUATIONRESULT
        case "notMatch":
            result = NOTMATCH_ASSIGNMENTFILTEREVALUATIONRESULT
        case "inconclusive":
            result = INCONCLUSIVE_ASSIGNMENTFILTEREVALUATIONRESULT
        case "failure":
            result = FAILURE_ASSIGNMENTFILTEREVALUATIONRESULT
        case "notEvaluated":
            result = NOTEVALUATED_ASSIGNMENTFILTEREVALUATIONRESULT
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAssignmentFilterEvaluationResult(values []AssignmentFilterEvaluationResult) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AssignmentFilterEvaluationResult) isMultiValue() bool {
    return false
}
