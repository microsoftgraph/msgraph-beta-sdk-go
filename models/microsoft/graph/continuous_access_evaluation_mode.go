package graph
import (
    "strings"
    "errors"
)
// 
type ContinuousAccessEvaluationMode int

const (
    STRICTENFORCEMENT_CONTINUOUSACCESSEVALUATIONMODE ContinuousAccessEvaluationMode = iota
    DISABLED_CONTINUOUSACCESSEVALUATIONMODE
    UNKNOWNFUTUREVALUE_CONTINUOUSACCESSEVALUATIONMODE
)

func (i ContinuousAccessEvaluationMode) String() string {
    return []string{"STRICTENFORCEMENT", "DISABLED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseContinuousAccessEvaluationMode(v string) (interface{}, error) {
    result := STRICTENFORCEMENT_CONTINUOUSACCESSEVALUATIONMODE
    switch strings.ToUpper(v) {
        case "STRICTENFORCEMENT":
            result = STRICTENFORCEMENT_CONTINUOUSACCESSEVALUATIONMODE
        case "DISABLED":
            result = DISABLED_CONTINUOUSACCESSEVALUATIONMODE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_CONTINUOUSACCESSEVALUATIONMODE
        default:
            return 0, errors.New("Unknown ContinuousAccessEvaluationMode value: " + v)
    }
    return &result, nil
}
func SerializeContinuousAccessEvaluationMode(values []ContinuousAccessEvaluationMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
