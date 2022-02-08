package graph
import (
    "strings"
    "errors"
)
// 
type AgentStatus int

const (
    ACTIVE_AGENTSTATUS AgentStatus = iota
    INACTIVE_AGENTSTATUS
)

func (i AgentStatus) String() string {
    return []string{"ACTIVE", "INACTIVE"}[i]
}
func ParseAgentStatus(v string) (interface{}, error) {
    result := ACTIVE_AGENTSTATUS
    switch strings.ToUpper(v) {
        case "ACTIVE":
            result = ACTIVE_AGENTSTATUS
        case "INACTIVE":
            result = INACTIVE_AGENTSTATUS
        default:
            return 0, errors.New("Unknown AgentStatus value: " + v)
    }
    return &result, nil
}
func SerializeAgentStatus(values []AgentStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
