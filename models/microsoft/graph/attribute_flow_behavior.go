package graph
import (
    "strings"
    "errors"
)
type AttributeFlowBehavior int

const (
    FLOWWHENCHANGED_ATTRIBUTEFLOWBEHAVIOR AttributeFlowBehavior = iota
    FLOWALWAYS_ATTRIBUTEFLOWBEHAVIOR
)

func (i AttributeFlowBehavior) String() string {
    return []string{"FLOWWHENCHANGED", "FLOWALWAYS"}[i]
}
func ParseAttributeFlowBehavior(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "FLOWWHENCHANGED":
            return FLOWWHENCHANGED_ATTRIBUTEFLOWBEHAVIOR, nil
        case "FLOWALWAYS":
            return FLOWALWAYS_ATTRIBUTEFLOWBEHAVIOR, nil
    }
    return 0, errors.New("Unknown AttributeFlowBehavior value: " + v)
}
func SerializeAttributeFlowBehavior(values []AttributeFlowBehavior) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
