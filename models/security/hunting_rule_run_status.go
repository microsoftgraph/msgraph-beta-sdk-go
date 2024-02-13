package security
import (
    "errors"
)
type HuntingRuleRunStatus int

const (
    RUNNING_HUNTINGRULERUNSTATUS HuntingRuleRunStatus = iota
    COMPLETED_HUNTINGRULERUNSTATUS
    FAILED_HUNTINGRULERUNSTATUS
    PARTIALLYFAILED_HUNTINGRULERUNSTATUS
    UNKNOWNFUTUREVALUE_HUNTINGRULERUNSTATUS
)

func (i HuntingRuleRunStatus) String() string {
    return []string{"running", "completed", "failed", "partiallyFailed", "unknownFutureValue"}[i]
}
func ParseHuntingRuleRunStatus(v string) (any, error) {
    result := RUNNING_HUNTINGRULERUNSTATUS
    switch v {
        case "running":
            result = RUNNING_HUNTINGRULERUNSTATUS
        case "completed":
            result = COMPLETED_HUNTINGRULERUNSTATUS
        case "failed":
            result = FAILED_HUNTINGRULERUNSTATUS
        case "partiallyFailed":
            result = PARTIALLYFAILED_HUNTINGRULERUNSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_HUNTINGRULERUNSTATUS
        default:
            return 0, errors.New("Unknown HuntingRuleRunStatus value: " + v)
    }
    return &result, nil
}
func SerializeHuntingRuleRunStatus(values []HuntingRuleRunStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i HuntingRuleRunStatus) isMultiValue() bool {
    return false
}
