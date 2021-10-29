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
    switch strings.ToUpper(v) {
        case "ALWAYS":
            return ALWAYS_ATTRIBUTEFLOWTYPE, nil
        case "OBJECTADDONLY":
            return OBJECTADDONLY_ATTRIBUTEFLOWTYPE, nil
        case "MULTIVALUEADDONLY":
            return MULTIVALUEADDONLY_ATTRIBUTEFLOWTYPE, nil
        case "VALUEADDONLY":
            return VALUEADDONLY_ATTRIBUTEFLOWTYPE, nil
        case "ATTRIBUTEADDONLY":
            return ATTRIBUTEADDONLY_ATTRIBUTEFLOWTYPE, nil
    }
    return 0, errors.New("Unknown AttributeFlowType value: " + v)
}
func SerializeAttributeFlowType(values []AttributeFlowType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
