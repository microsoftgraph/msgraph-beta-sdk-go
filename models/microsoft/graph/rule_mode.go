package graph
import (
    "strings"
    "errors"
)
// 
type RuleMode int

const (
    AUDIT_RULEMODE RuleMode = iota
    AUDITANDNOTIFY_RULEMODE
    ENFORCE_RULEMODE
    PENDINGDELETION_RULEMODE
    TEST_RULEMODE
)

func (i RuleMode) String() string {
    return []string{"AUDIT", "AUDITANDNOTIFY", "ENFORCE", "PENDINGDELETION", "TEST"}[i]
}
func ParseRuleMode(v string) (interface{}, error) {
    result := AUDIT_RULEMODE
    switch strings.ToUpper(v) {
        case "AUDIT":
            result = AUDIT_RULEMODE
        case "AUDITANDNOTIFY":
            result = AUDITANDNOTIFY_RULEMODE
        case "ENFORCE":
            result = ENFORCE_RULEMODE
        case "PENDINGDELETION":
            result = PENDINGDELETION_RULEMODE
        case "TEST":
            result = TEST_RULEMODE
        default:
            return 0, errors.New("Unknown RuleMode value: " + v)
    }
    return &result, nil
}
func SerializeRuleMode(values []RuleMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
