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
    switch strings.ToUpper(v) {
        case "STRICTENFORCEMENT":
            return STRICTENFORCEMENT_CONTINUOUSACCESSEVALUATIONMODE, nil
        case "DISABLED":
            return DISABLED_CONTINUOUSACCESSEVALUATIONMODE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_CONTINUOUSACCESSEVALUATIONMODE, nil
    }
    return 0, errors.New("Unknown ContinuousAccessEvaluationMode value: " + v)
}
func SerializeContinuousAccessEvaluationMode(values []ContinuousAccessEvaluationMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
