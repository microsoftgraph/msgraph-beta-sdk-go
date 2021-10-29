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
    switch strings.ToUpper(v) {
        case "ACTIVE":
            return ACTIVE_AGENTSTATUS, nil
        case "INACTIVE":
            return INACTIVE_AGENTSTATUS, nil
    }
    return 0, errors.New("Unknown AgentStatus value: " + v)
}
func SerializeAgentStatus(values []AgentStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
