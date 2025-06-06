// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package agentic
type AgentType int

const (
    NOTAGENTIC_AGENTTYPE AgentType = iota
    AGENTICAPPBUILDER_AGENTTYPE
    AGENTICAPP_AGENTTYPE
    AGENTICAPPINSTANCE_AGENTTYPE
    UNKNOWNFUTUREVALUE_AGENTTYPE
)

func (i AgentType) String() string {
    return []string{"notAgentic", "agenticAppBuilder", "agenticApp", "agenticAppInstance", "unknownFutureValue"}[i]
}
func ParseAgentType(v string) (any, error) {
    result := NOTAGENTIC_AGENTTYPE
    switch v {
        case "notAgentic":
            result = NOTAGENTIC_AGENTTYPE
        case "agenticAppBuilder":
            result = AGENTICAPPBUILDER_AGENTTYPE
        case "agenticApp":
            result = AGENTICAPP_AGENTTYPE
        case "agenticAppInstance":
            result = AGENTICAPPINSTANCE_AGENTTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_AGENTTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAgentType(values []AgentType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AgentType) isMultiValue() bool {
    return false
}
