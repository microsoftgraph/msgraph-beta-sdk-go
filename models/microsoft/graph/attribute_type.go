package graph
import (
    "strings"
    "errors"
)
// 
type AttributeType int

const (
    STRING_ATTRIBUTETYPE AttributeType = iota
    INTEGER_ATTRIBUTETYPE
    REFERENCE_ATTRIBUTETYPE
    BINARY_ATTRIBUTETYPE
    BOOLEAN_ATTRIBUTETYPE
    DATETIME_ATTRIBUTETYPE
)

func (i AttributeType) String() string {
    return []string{"STRING", "INTEGER", "REFERENCE", "BINARY", "BOOLEAN", "DATETIME"}[i]
}
func ParseAttributeType(v string) (interface{}, error) {
    result := STRING_ATTRIBUTETYPE
    switch strings.ToUpper(v) {
        case "STRING":
            result = STRING_ATTRIBUTETYPE
        case "INTEGER":
            result = INTEGER_ATTRIBUTETYPE
        case "REFERENCE":
            result = REFERENCE_ATTRIBUTETYPE
        case "BINARY":
            result = BINARY_ATTRIBUTETYPE
        case "BOOLEAN":
            result = BOOLEAN_ATTRIBUTETYPE
        case "DATETIME":
            result = DATETIME_ATTRIBUTETYPE
        default:
            return 0, errors.New("Unknown AttributeType value: " + v)
    }
    return &result, nil
}
func SerializeAttributeType(values []AttributeType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
