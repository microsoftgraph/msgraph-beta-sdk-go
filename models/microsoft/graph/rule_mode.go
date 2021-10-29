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
    switch strings.ToUpper(v) {
        case "AUDIT":
            return AUDIT_RULEMODE, nil
        case "AUDITANDNOTIFY":
            return AUDITANDNOTIFY_RULEMODE, nil
        case "ENFORCE":
            return ENFORCE_RULEMODE, nil
        case "PENDINGDELETION":
            return PENDINGDELETION_RULEMODE, nil
        case "TEST":
            return TEST_RULEMODE, nil
    }
    return 0, errors.New("Unknown RuleMode value: " + v)
}
func SerializeRuleMode(values []RuleMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
