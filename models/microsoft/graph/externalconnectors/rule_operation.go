package externalconnectors
import (
    "strings"
    "errors"
)
// 
type RuleOperation int

const (
    EQUALS_RULEOPERATION RuleOperation = iota
    NOTEQUALS_RULEOPERATION
    CONTAINS_RULEOPERATION
    NOTCONTAINS_RULEOPERATION
    LESSTHAN_RULEOPERATION
    GREATERTHAN_RULEOPERATION
    STARTSWITH_RULEOPERATION
    UNKNOWNFUTUREVALUE_RULEOPERATION
)

func (i RuleOperation) String() string {
    return []string{"EQUALS", "NOTEQUALS", "CONTAINS", "NOTCONTAINS", "LESSTHAN", "GREATERTHAN", "STARTSWITH", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseRuleOperation(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "EQUALS":
            return EQUALS_RULEOPERATION, nil
        case "NOTEQUALS":
            return NOTEQUALS_RULEOPERATION, nil
        case "CONTAINS":
            return CONTAINS_RULEOPERATION, nil
        case "NOTCONTAINS":
            return NOTCONTAINS_RULEOPERATION, nil
        case "LESSTHAN":
            return LESSTHAN_RULEOPERATION, nil
        case "GREATERTHAN":
            return GREATERTHAN_RULEOPERATION, nil
        case "STARTSWITH":
            return STARTSWITH_RULEOPERATION, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_RULEOPERATION, nil
    }
    return 0, errors.New("Unknown RuleOperation value: " + v)
}
func SerializeRuleOperation(values []RuleOperation) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
