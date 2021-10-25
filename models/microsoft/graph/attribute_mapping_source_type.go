package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "ATTRIBUTE":
            return ATTRIBUTE_ATTRIBUTEMAPPINGSOURCETYPE, nil
        case "CONSTANT":
            return CONSTANT_ATTRIBUTEMAPPINGSOURCETYPE, nil
        case "FUNCTION":
            return FUNCTION_ATTRIBUTEMAPPINGSOURCETYPE, nil
    }
    return 0, errors.New("Unknown AttributeMappingSourceType value: " + v)
}
func SerializeAttributeMappingSourceType(values []AttributeMappingSourceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
