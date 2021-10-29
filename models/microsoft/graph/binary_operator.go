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
    switch strings.ToUpper(v) {
        case "OR":
            return OR_BINARYOPERATOR, nil
        case "AND":
            return AND_BINARYOPERATOR, nil
    }
    return 0, errors.New("Unknown BinaryOperator value: " + v)
}
func SerializeBinaryOperator(values []BinaryOperator) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
