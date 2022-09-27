package models
import (
    "errors"
)
// Provides operations to manage the collection of accessReviewDecision entities.
type AttributeFlowType int

const (
    ALWAYS_ATTRIBUTEFLOWTYPE AttributeFlowType = iota
    OBJECTADDONLY_ATTRIBUTEFLOWTYPE
    MULTIVALUEADDONLY_ATTRIBUTEFLOWTYPE
    VALUEADDONLY_ATTRIBUTEFLOWTYPE
    ATTRIBUTEADDONLY_ATTRIBUTEFLOWTYPE
)

func (i AttributeFlowType) String() string {
    return []string{"Always", "ObjectAddOnly", "MultiValueAddOnly", "ValueAddOnly", "AttributeAddOnly"}[i]
}
func ParseAttributeFlowType(v string) (interface{}, error) {
    result := ALWAYS_ATTRIBUTEFLOWTYPE
    switch v {
        case "Always":
            result = ALWAYS_ATTRIBUTEFLOWTYPE
        case "ObjectAddOnly":
            result = OBJECTADDONLY_ATTRIBUTEFLOWTYPE
        case "MultiValueAddOnly":
            result = MULTIVALUEADDONLY_ATTRIBUTEFLOWTYPE
        case "ValueAddOnly":
            result = VALUEADDONLY_ATTRIBUTEFLOWTYPE
        case "AttributeAddOnly":
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
