package graph
import (
    "strings"
    "errors"
)
// 
type ClassificationMethod int

const (
    PATTERNMATCH_CLASSIFICATIONMETHOD ClassificationMethod = iota
    EXACTDATAMATCH_CLASSIFICATIONMETHOD
    FINGERPRINT_CLASSIFICATIONMETHOD
    MACHINELEARNING_CLASSIFICATIONMETHOD
)

func (i ClassificationMethod) String() string {
    return []string{"PATTERNMATCH", "EXACTDATAMATCH", "FINGERPRINT", "MACHINELEARNING"}[i]
}
func ParseClassificationMethod(v string) (interface{}, error) {
    result := PATTERNMATCH_CLASSIFICATIONMETHOD
    switch strings.ToUpper(v) {
        case "PATTERNMATCH":
            result = PATTERNMATCH_CLASSIFICATIONMETHOD
        case "EXACTDATAMATCH":
            result = EXACTDATAMATCH_CLASSIFICATIONMETHOD
        case "FINGERPRINT":
            result = FINGERPRINT_CLASSIFICATIONMETHOD
        case "MACHINELEARNING":
            result = MACHINELEARNING_CLASSIFICATIONMETHOD
        default:
            return 0, errors.New("Unknown ClassificationMethod value: " + v)
    }
    return &result, nil
}
func SerializeClassificationMethod(values []ClassificationMethod) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
