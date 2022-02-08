package graph
import (
    "strings"
    "errors"
)
// 
type AttributeMappingSourceType int

const (
    ATTRIBUTE_ATTRIBUTEMAPPINGSOURCETYPE AttributeMappingSourceType = iota
    CONSTANT_ATTRIBUTEMAPPINGSOURCETYPE
    FUNCTION_ATTRIBUTEMAPPINGSOURCETYPE
)

func (i AttributeMappingSourceType) String() string {
    return []string{"ATTRIBUTE", "CONSTANT", "FUNCTION"}[i]
}
func ParseAttributeMappingSourceType(v string) (interface{}, error) {
    result := ATTRIBUTE_ATTRIBUTEMAPPINGSOURCETYPE
    switch strings.ToUpper(v) {
        case "ATTRIBUTE":
            result = ATTRIBUTE_ATTRIBUTEMAPPINGSOURCETYPE
        case "CONSTANT":
            result = CONSTANT_ATTRIBUTEMAPPINGSOURCETYPE
        case "FUNCTION":
            result = FUNCTION_ATTRIBUTEMAPPINGSOURCETYPE
        default:
            return 0, errors.New("Unknown AttributeMappingSourceType value: " + v)
    }
    return &result, nil
}
func SerializeAttributeMappingSourceType(values []AttributeMappingSourceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
