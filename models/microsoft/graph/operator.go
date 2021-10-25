package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_OPERATOR, nil
        case "AND":
            return AND_OPERATOR, nil
        case "OR":
            return OR_OPERATOR, nil
        case "ISEQUALS":
            return ISEQUALS_OPERATOR, nil
        case "NOTEQUALS":
            return NOTEQUALS_OPERATOR, nil
        case "GREATERTHAN":
            return GREATERTHAN_OPERATOR, nil
        case "LESSTHAN":
            return LESSTHAN_OPERATOR, nil
        case "BETWEEN":
            return BETWEEN_OPERATOR, nil
        case "NOTBETWEEN":
            return NOTBETWEEN_OPERATOR, nil
        case "GREATEREQUALS":
            return GREATEREQUALS_OPERATOR, nil
        case "LESSEQUALS":
            return LESSEQUALS_OPERATOR, nil
        case "DAYTIMEBETWEEN":
            return DAYTIMEBETWEEN_OPERATOR, nil
        case "BEGINSWITH":
            return BEGINSWITH_OPERATOR, nil
        case "NOTBEGINSWITH":
            return NOTBEGINSWITH_OPERATOR, nil
        case "ENDSWITH":
            return ENDSWITH_OPERATOR, nil
        case "NOTENDSWITH":
            return NOTENDSWITH_OPERATOR, nil
        case "CONTAINS":
            return CONTAINS_OPERATOR, nil
        case "NOTCONTAINS":
            return NOTCONTAINS_OPERATOR, nil
        case "ALLOF":
            return ALLOF_OPERATOR, nil
        case "ONEOF":
            return ONEOF_OPERATOR, nil
        case "NONEOF":
            return NONEOF_OPERATOR, nil
        case "SETEQUALS":
            return SETEQUALS_OPERATOR, nil
        case "ORDEREDSETEQUALS":
            return ORDEREDSETEQUALS_OPERATOR, nil
        case "SUBSETOF":
            return SUBSETOF_OPERATOR, nil
        case "EXCLUDESALL":
            return EXCLUDESALL_OPERATOR, nil
    }
    return 0, errors.New("Unknown Operator value: " + v)
}
func SerializeOperator(values []Operator) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
