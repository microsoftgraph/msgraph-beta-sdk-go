package graph
import (
    "strings"
    "errors"
)
// 
type AttributeFlowType int

const (
    ALWAYS_ATTRIBUTEFLOWTYPE AttributeFlowType = iota
    OBJECTADDONLY_ATTRIBUTEFLOWTYPE
    MULTIVALUEADDONLY_ATTRIBUTEFLOWTYPE
    VALUEADDONLY_ATTRIBUTEFLOWTYPE
    ATTRIBUTEADDONLY_ATTRIBUTEFLOWTYPE
)

func (i AttributeFlowType) String() string {
    return []string{"ALWAYS", "OBJECTADDONLY", "MULTIVALUEADDONLY", "VALUEADDONLY", "ATTRIBUTEADDONLY"}[i]
}
func ParseAttributeFlowType(v string) (interface{}, error) {
    result := ALWAYS_ATTRIBUTEFLOWTYPE
    switch strings.ToUpper(v) {
        case "ALWAYS":
            result = ALWAYS_ATTRIBUTEFLOWTYPE
        case "OBJECTADDONLY":
            result = OBJECTADDONLY_ATTRIBUTEFLOWTYPE
        case "MULTIVALUEADDONLY":
            result = MULTIVALUEADDONLY_ATTRIBUTEFLOWTYPE
        case "VALUEADDONLY":
            result = VALUEADDONLY_ATTRIBUTEFLOWTYPE
        case "ATTRIBUTEADDONLY":
            result = ATTRIBUTEADDONLY_ATTRIBUTEFLOWTYPE
        default:
            return 0, errors.New("Unknown AttributeFlowType value: " + v)
    }
    return &result, nil
}
func SerializeAttributeFlowType(values []AttributeFlowType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
