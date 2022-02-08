package graph
import (
    "strings"
    "errors"
)
// 
type Operator int

const (
    NONE_OPERATOR Operator = iota
    AND_OPERATOR
    OR_OPERATOR
    ISEQUALS_OPERATOR
    NOTEQUALS_OPERATOR
    GREATERTHAN_OPERATOR
    LESSTHAN_OPERATOR
    BETWEEN_OPERATOR
    NOTBETWEEN_OPERATOR
    GREATEREQUALS_OPERATOR
    LESSEQUALS_OPERATOR
    DAYTIMEBETWEEN_OPERATOR
    BEGINSWITH_OPERATOR
    NOTBEGINSWITH_OPERATOR
    ENDSWITH_OPERATOR
    NOTENDSWITH_OPERATOR
    CONTAINS_OPERATOR
    NOTCONTAINS_OPERATOR
    ALLOF_OPERATOR
    ONEOF_OPERATOR
    NONEOF_OPERATOR
    SETEQUALS_OPERATOR
    ORDEREDSETEQUALS_OPERATOR
    SUBSETOF_OPERATOR
    EXCLUDESALL_OPERATOR
)

func (i Operator) String() string {
    return []string{"NONE", "AND", "OR", "ISEQUALS", "NOTEQUALS", "GREATERTHAN", "LESSTHAN", "BETWEEN", "NOTBETWEEN", "GREATEREQUALS", "LESSEQUALS", "DAYTIMEBETWEEN", "BEGINSWITH", "NOTBEGINSWITH", "ENDSWITH", "NOTENDSWITH", "CONTAINS", "NOTCONTAINS", "ALLOF", "ONEOF", "NONEOF", "SETEQUALS", "ORDEREDSETEQUALS", "SUBSETOF", "EXCLUDESALL"}[i]
}
func ParseOperator(v string) (interface{}, error) {
    result := NONE_OPERATOR
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_OPERATOR
        case "AND":
            result = AND_OPERATOR
        case "OR":
            result = OR_OPERATOR
        case "ISEQUALS":
            result = ISEQUALS_OPERATOR
        case "NOTEQUALS":
            result = NOTEQUALS_OPERATOR
        case "GREATERTHAN":
            result = GREATERTHAN_OPERATOR
        case "LESSTHAN":
            result = LESSTHAN_OPERATOR
        case "BETWEEN":
            result = BETWEEN_OPERATOR
        case "NOTBETWEEN":
            result = NOTBETWEEN_OPERATOR
        case "GREATEREQUALS":
            result = GREATEREQUALS_OPERATOR
        case "LESSEQUALS":
            result = LESSEQUALS_OPERATOR
        case "DAYTIMEBETWEEN":
            result = DAYTIMEBETWEEN_OPERATOR
        case "BEGINSWITH":
            result = BEGINSWITH_OPERATOR
        case "NOTBEGINSWITH":
            result = NOTBEGINSWITH_OPERATOR
        case "ENDSWITH":
            result = ENDSWITH_OPERATOR
        case "NOTENDSWITH":
            result = NOTENDSWITH_OPERATOR
        case "CONTAINS":
            result = CONTAINS_OPERATOR
        case "NOTCONTAINS":
            result = NOTCONTAINS_OPERATOR
        case "ALLOF":
            result = ALLOF_OPERATOR
        case "ONEOF":
            result = ONEOF_OPERATOR
        case "NONEOF":
            result = NONEOF_OPERATOR
        case "SETEQUALS":
            result = SETEQUALS_OPERATOR
        case "ORDEREDSETEQUALS":
            result = ORDEREDSETEQUALS_OPERATOR
        case "SUBSETOF":
            result = SUBSETOF_OPERATOR
        case "EXCLUDESALL":
            result = EXCLUDESALL_OPERATOR
        default:
            return 0, errors.New("Unknown Operator value: " + v)
    }
    return &result, nil
}
func SerializeOperator(values []Operator) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
