package models
// Represents the different type of operators which can be used to craft the AssignmentFilter rule.
type AssignmentFilterOperator int

const (
    // Indicates operator is not set
    NOTSET_ASSIGNMENTFILTEROPERATOR AssignmentFilterOperator = iota
    // Indicates the devices whose property value equals the configured input in Assignment Filters.
    EQUALS_ASSIGNMENTFILTEROPERATOR
    // Indicates the devices whose property value not equals the configured input in Assignment Filters.
    NOTEQUALS_ASSIGNMENTFILTEROPERATOR
    // Indicates the devices whose property value starts with the configured input in Assignment Filters.
    STARTSWITH_ASSIGNMENTFILTEROPERATOR
    // Indicates the devices whose property value does not start with the configured input in Assignment Filters.
    NOTSTARTSWITH_ASSIGNMENTFILTEROPERATOR
    // Indicates the devices whose property value contains the configured input in Assignment Filters.
    CONTAINS_ASSIGNMENTFILTEROPERATOR
    // Indicates the devices whose property value does not contain the configured input in Assignment Filters.
    NOTCONTAINS_ASSIGNMENTFILTEROPERATOR
    // Indicates the devices whose property value equals one of the configured input in Assignment Filters.
    IN_ASSIGNMENTFILTEROPERATOR
    // Indicates the devices whose property value does not equals any of the configured input in Assignment Filters.
    NOTIN_ASSIGNMENTFILTEROPERATOR
    // Indicates the devices whose property value ends with the configured input in Assignment Filters.
    ENDSWITH_ASSIGNMENTFILTEROPERATOR
    // Indicates the devices whose property value does not end with the configured input in Assignment Filters.
    NOTENDSWITH_ASSIGNMENTFILTEROPERATOR
    // Indicates the devices whose property value is greater than the configured input in Assignment Filters.
    GREATERTHAN_ASSIGNMENTFILTEROPERATOR
    // `Indicates the devices whose property value is greater than or equal to the configured input in Assignment Filters.
    GREATERTHANOREQUALS_ASSIGNMENTFILTEROPERATOR
    // Indicates the devices whose property value is less than the configured input in Assignment Filters.
    LESSTHAN_ASSIGNMENTFILTEROPERATOR
    // Indicates the devices whose property value is less than or equal to the configured input in Assignment Filters.
    LESSTHANOREQUALS_ASSIGNMENTFILTEROPERATOR
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_ASSIGNMENTFILTEROPERATOR
)

func (i AssignmentFilterOperator) String() string {
    return []string{"notSet", "equals", "notEquals", "startsWith", "notStartsWith", "contains", "notContains", "in", "notIn", "endsWith", "notEndsWith", "greaterThan", "greaterThanOrEquals", "lessThan", "lessThanOrEquals", "unknownFutureValue"}[i]
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
        case "greaterThan":
            result = GREATERTHAN_ASSIGNMENTFILTEROPERATOR
        case "greaterThanOrEquals":
            result = GREATERTHANOREQUALS_ASSIGNMENTFILTEROPERATOR
        case "lessThan":
            result = LESSTHAN_ASSIGNMENTFILTEROPERATOR
        case "lessThanOrEquals":
            result = LESSTHANOREQUALS_ASSIGNMENTFILTEROPERATOR
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ASSIGNMENTFILTEROPERATOR
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
