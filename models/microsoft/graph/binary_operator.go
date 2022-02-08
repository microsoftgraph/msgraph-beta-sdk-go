package graph
import (
    "strings"
    "errors"
)
// 
type BinaryOperator int

const (
    OR_BINARYOPERATOR BinaryOperator = iota
    AND_BINARYOPERATOR
)

func (i BinaryOperator) String() string {
    return []string{"OR", "AND"}[i]
}
func ParseBinaryOperator(v string) (interface{}, error) {
    result := OR_BINARYOPERATOR
    switch strings.ToUpper(v) {
        case "OR":
            result = OR_BINARYOPERATOR
        case "AND":
            result = AND_BINARYOPERATOR
        default:
            return 0, errors.New("Unknown BinaryOperator value: " + v)
    }
    return &result, nil
}
func SerializeBinaryOperator(values []BinaryOperator) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
