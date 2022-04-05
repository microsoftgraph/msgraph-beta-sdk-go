package models
import (
    "strings"
    "errors"
)
// Provides operations to call the getSupportedProperties method.
type AssignmentFilterOperator int

const (
    NOTSET_ASSIGNMENTFILTEROPERATOR AssignmentFilterOperator = iota
    EQUALS_ASSIGNMENTFILTEROPERATOR
    NOTEQUALS_ASSIGNMENTFILTEROPERATOR
    STARTSWITH_ASSIGNMENTFILTEROPERATOR
    NOTSTARTSWITH_ASSIGNMENTFILTEROPERATOR
    CONTAINS_ASSIGNMENTFILTEROPERATOR
    NOTCONTAINS_ASSIGNMENTFILTEROPERATOR
    IN_ASSIGNMENTFILTEROPERATOR
    NOTIN_ASSIGNMENTFILTEROPERATOR
    ENDSWITH_ASSIGNMENTFILTEROPERATOR
    NOTENDSWITH_ASSIGNMENTFILTEROPERATOR
)

func (i AssignmentFilterOperator) String() string {
    return []string{"NOTSET", "EQUALS", "NOTEQUALS", "STARTSWITH", "NOTSTARTSWITH", "CONTAINS", "NOTCONTAINS", "IN", "NOTIN", "ENDSWITH", "NOTENDSWITH"}[i]
}
func ParseAssignmentFilterOperator(v string) (interface{}, error) {
    result := NOTSET_ASSIGNMENTFILTEROPERATOR
    switch strings.ToUpper(v) {
        case "NOTSET":
            result = NOTSET_ASSIGNMENTFILTEROPERATOR
        case "EQUALS":
            result = EQUALS_ASSIGNMENTFILTEROPERATOR
        case "NOTEQUALS":
            result = NOTEQUALS_ASSIGNMENTFILTEROPERATOR
        case "STARTSWITH":
            result = STARTSWITH_ASSIGNMENTFILTEROPERATOR
        case "NOTSTARTSWITH":
            result = NOTSTARTSWITH_ASSIGNMENTFILTEROPERATOR
        case "CONTAINS":
            result = CONTAINS_ASSIGNMENTFILTEROPERATOR
        case "NOTCONTAINS":
            result = NOTCONTAINS_ASSIGNMENTFILTEROPERATOR
        case "IN":
            result = IN_ASSIGNMENTFILTEROPERATOR
        case "NOTIN":
            result = NOTIN_ASSIGNMENTFILTEROPERATOR
        case "ENDSWITH":
            result = ENDSWITH_ASSIGNMENTFILTEROPERATOR
        case "NOTENDSWITH":
            result = NOTENDSWITH_ASSIGNMENTFILTEROPERATOR
        default:
            return 0, errors.New("Unknown AssignmentFilterOperator value: " + v)
    }
    return &result, nil
}
func SerializeAssignmentFilterOperator(values []AssignmentFilterOperator) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
