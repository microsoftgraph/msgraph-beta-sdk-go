package graph
import (
    "strings"
    "errors"
)
// 
type AssignmentFilterEvaluationResult int

const (
    UNKNOWN_ASSIGNMENTFILTEREVALUATIONRESULT AssignmentFilterEvaluationResult = iota
    MATCH_ASSIGNMENTFILTEREVALUATIONRESULT
    NOTMATCH_ASSIGNMENTFILTEREVALUATIONRESULT
    INCONCLUSIVE_ASSIGNMENTFILTEREVALUATIONRESULT
    FAILURE_ASSIGNMENTFILTEREVALUATIONRESULT
    NOTEVALUATED_ASSIGNMENTFILTEREVALUATIONRESULT
)

func (i AssignmentFilterEvaluationResult) String() string {
    return []string{"UNKNOWN", "MATCH", "NOTMATCH", "INCONCLUSIVE", "FAILURE", "NOTEVALUATED"}[i]
}
func ParseAssignmentFilterEvaluationResult(v string) (interface{}, error) {
    result := UNKNOWN_ASSIGNMENTFILTEREVALUATIONRESULT
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_ASSIGNMENTFILTEREVALUATIONRESULT
        case "MATCH":
            result = MATCH_ASSIGNMENTFILTEREVALUATIONRESULT
        case "NOTMATCH":
            result = NOTMATCH_ASSIGNMENTFILTEREVALUATIONRESULT
        case "INCONCLUSIVE":
            result = INCONCLUSIVE_ASSIGNMENTFILTEREVALUATIONRESULT
        case "FAILURE":
            result = FAILURE_ASSIGNMENTFILTEREVALUATIONRESULT
        case "NOTEVALUATED":
            result = NOTEVALUATED_ASSIGNMENTFILTEREVALUATIONRESULT
        default:
            return 0, errors.New("Unknown AssignmentFilterEvaluationResult value: " + v)
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
