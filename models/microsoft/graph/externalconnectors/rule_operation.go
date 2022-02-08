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
    result := EQUALS_RULEOPERATION
    switch strings.ToUpper(v) {
        case "EQUALS":
            result = EQUALS_RULEOPERATION
        case "NOTEQUALS":
            result = NOTEQUALS_RULEOPERATION
        case "CONTAINS":
            result = CONTAINS_RULEOPERATION
        case "NOTCONTAINS":
            result = NOTCONTAINS_RULEOPERATION
        case "LESSTHAN":
            result = LESSTHAN_RULEOPERATION
        case "GREATERTHAN":
            result = GREATERTHAN_RULEOPERATION
        case "STARTSWITH":
            result = STARTSWITH_RULEOPERATION
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_RULEOPERATION
        default:
            return 0, errors.New("Unknown RuleOperation value: " + v)
    }
    return &result, nil
}
func SerializeRuleOperation(values []RuleOperation) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
