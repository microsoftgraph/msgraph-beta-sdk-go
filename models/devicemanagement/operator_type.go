package devicemanagement
import (
    "errors"
)
// 
type OperatorType int

const (
    GREATEROREQUAL_OPERATORTYPE OperatorType = iota
    EQUAL_OPERATORTYPE
    GREATER_OPERATORTYPE
    LESS_OPERATORTYPE
    LESSOREQUAL_OPERATORTYPE
    NOTEQUAL_OPERATORTYPE
    UNKNOWNFUTUREVALUE_OPERATORTYPE
)

func (i OperatorType) String() string {
    return []string{"greaterOrEqual", "equal", "greater", "less", "lessOrEqual", "notEqual", "unknownFutureValue"}[i]
}
func ParseOperatorType(v string) (any, error) {
    result := GREATEROREQUAL_OPERATORTYPE
    switch v {
        case "greaterOrEqual":
            result = GREATEROREQUAL_OPERATORTYPE
        case "equal":
            result = EQUAL_OPERATORTYPE
        case "greater":
            result = GREATER_OPERATORTYPE
        case "less":
            result = LESS_OPERATORTYPE
        case "lessOrEqual":
            result = LESSOREQUAL_OPERATORTYPE
        case "notEqual":
            result = NOTEQUAL_OPERATORTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_OPERATORTYPE
        default:
            return 0, errors.New("Unknown OperatorType value: " + v)
    }
    return &result, nil
}
func SerializeOperatorType(values []OperatorType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OperatorType) isMultiValue() bool {
    return false
}
