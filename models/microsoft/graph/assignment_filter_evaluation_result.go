package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_ASSIGNMENTFILTEREVALUATIONRESULT, nil
        case "MATCH":
            return MATCH_ASSIGNMENTFILTEREVALUATIONRESULT, nil
        case "NOTMATCH":
            return NOTMATCH_ASSIGNMENTFILTEREVALUATIONRESULT, nil
        case "INCONCLUSIVE":
            return INCONCLUSIVE_ASSIGNMENTFILTEREVALUATIONRESULT, nil
        case "FAILURE":
            return FAILURE_ASSIGNMENTFILTEREVALUATIONRESULT, nil
        case "NOTEVALUATED":
            return NOTEVALUATED_ASSIGNMENTFILTEREVALUATIONRESULT, nil
    }
    return 0, errors.New("Unknown AssignmentFilterEvaluationResult value: " + v)
}
func SerializeAssignmentFilterEvaluationResult(values []AssignmentFilterEvaluationResult) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
