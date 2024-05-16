package models
// Represents the different type of operators which can be used to craft the AssignmentFilter rule.
type AssignmentFilterOperator int

const (
    // NotSet.
    NOTSET_ASSIGNMENTFILTEROPERATOR AssignmentFilterOperator = iota
    // Equals.
    EQUALS_ASSIGNMENTFILTEROPERATOR
    // NotEquals.
    NOTEQUALS_ASSIGNMENTFILTEROPERATOR
    // StartsWith.
    STARTSWITH_ASSIGNMENTFILTEROPERATOR
    // NotStartsWith.
    NOTSTARTSWITH_ASSIGNMENTFILTEROPERATOR
    // Contains.
    CONTAINS_ASSIGNMENTFILTEROPERATOR
    // NotContains.
    NOTCONTAINS_ASSIGNMENTFILTEROPERATOR
    // In.
    IN_ASSIGNMENTFILTEROPERATOR
    // NotIn.
    NOTIN_ASSIGNMENTFILTEROPERATOR
    // EndsWith.
    ENDSWITH_ASSIGNMENTFILTEROPERATOR
    // NotEndsWith.
    NOTENDSWITH_ASSIGNMENTFILTEROPERATOR
)

func (i AssignmentFilterOperator) String() string {
    return []string{"notSet", "equals", "notEquals", "startsWith", "notStartsWith", "contains", "notContains", "in", "notIn", "endsWith", "notEndsWith"}[i]
}
func ParseAssignmentFilterOperator(v string) (any, error) {
    result := NOTSET_ASSIGNMENTFILTEROPERATOR
    switch v {
        case "notSet":
            result = NOTSET_ASSIGNMENTFILTEROPERATOR
        case "equals":
            result = EQUALS_ASSIGNMENTFILTEROPERATOR
        case "notEquals":
            result = NOTEQUALS_ASSIGNMENTFILTEROPERATOR
        case "startsWith":
            result = STARTSWITH_ASSIGNMENTFILTEROPERATOR
        case "notStartsWith":
            result = NOTSTARTSWITH_ASSIGNMENTFILTEROPERATOR
        case "contains":
            result = CONTAINS_ASSIGNMENTFILTEROPERATOR
        case "notContains":
            result = NOTCONTAINS_ASSIGNMENTFILTEROPERATOR
        case "in":
            result = IN_ASSIGNMENTFILTEROPERATOR
        case "notIn":
            result = NOTIN_ASSIGNMENTFILTEROPERATOR
        case "endsWith":
            result = ENDSWITH_ASSIGNMENTFILTEROPERATOR
        case "notEndsWith":
            result = NOTENDSWITH_ASSIGNMENTFILTEROPERATOR
        default:
            return nil, nil
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
func (i AssignmentFilterOperator) isMultiValue() bool {
    return false
}
