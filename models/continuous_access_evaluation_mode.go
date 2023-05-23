package models
import (
    "errors"
)
// 
type ContinuousAccessEvaluationMode int

const (
    STRICTENFORCEMENT_CONTINUOUSACCESSEVALUATIONMODE ContinuousAccessEvaluationMode = iota
    DISABLED_CONTINUOUSACCESSEVALUATIONMODE
    UNKNOWNFUTUREVALUE_CONTINUOUSACCESSEVALUATIONMODE
    STRICTLOCATION_CONTINUOUSACCESSEVALUATIONMODE
)

func (i ContinuousAccessEvaluationMode) String() string {
    return []string{"strictEnforcement", "disabled", "unknownFutureValue", "strictLocation"}[i]
}
func ParseContinuousAccessEvaluationMode(v string) (any, error) {
    result := STRICTENFORCEMENT_CONTINUOUSACCESSEVALUATIONMODE
    switch v {
        case "strictEnforcement":
            result = STRICTENFORCEMENT_CONTINUOUSACCESSEVALUATIONMODE
        case "disabled":
            result = DISABLED_CONTINUOUSACCESSEVALUATIONMODE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CONTINUOUSACCESSEVALUATIONMODE
        case "strictLocation":
            result = STRICTLOCATION_CONTINUOUSACCESSEVALUATIONMODE
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
