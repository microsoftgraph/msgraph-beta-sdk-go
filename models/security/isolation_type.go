package security
import (
    "errors"
)
// 
type IsolationType int

const (
    FULL_ISOLATIONTYPE IsolationType = iota
    SELECTIVE_ISOLATIONTYPE
    UNKNOWNFUTUREVALUE_ISOLATIONTYPE
)

func (i IsolationType) String() string {
    return []string{"full", "selective", "unknownFutureValue"}[i]
}
func ParseIsolationType(v string) (any, error) {
    result := FULL_ISOLATIONTYPE
    switch v {
        case "full":
            result = FULL_ISOLATIONTYPE
        case "selective":
            result = SELECTIVE_ISOLATIONTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ISOLATIONTYPE
        default:
            return 0, errors.New("Unknown IsolationType value: " + v)
    }
    return &result, nil
}
func SerializeIsolationType(values []IsolationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i IsolationType) isMultiValue() bool {
    return false
}
