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
    switch strings.ToUpper(v) {
        case "STRING":
            return STRING_ATTRIBUTETYPE, nil
        case "INTEGER":
            return INTEGER_ATTRIBUTETYPE, nil
        case "REFERENCE":
            return REFERENCE_ATTRIBUTETYPE, nil
        case "BINARY":
            return BINARY_ATTRIBUTETYPE, nil
        case "BOOLEAN":
            return BOOLEAN_ATTRIBUTETYPE, nil
        case "DATETIME":
            return DATETIME_ATTRIBUTETYPE, nil
    }
    return 0, errors.New("Unknown AttributeType value: " + v)
}
func SerializeAttributeType(values []AttributeType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
