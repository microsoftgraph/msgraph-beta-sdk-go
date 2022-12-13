package models
import (
    "errors"
)
// Casts the previous resource to application.
type WeakAlgorithms int

const (
    RSASHA1_WEAKALGORITHMS WeakAlgorithms = iota
    UNKNOWNFUTUREVALUE_WEAKALGORITHMS
)

func (i WeakAlgorithms) String() string {
    return []string{"rsaSha1", "unknownFutureValue"}[i]
}
func ParseWeakAlgorithms(v string) (interface{}, error) {
    result := RSASHA1_WEAKALGORITHMS
    switch v {
        case "rsaSha1":
            result = RSASHA1_WEAKALGORITHMS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_WEAKALGORITHMS
        default:
            return 0, errors.New("Unknown WeakAlgorithms value: " + v)
    }
    return &result, nil
}
func SerializeWeakAlgorithms(values []WeakAlgorithms) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
