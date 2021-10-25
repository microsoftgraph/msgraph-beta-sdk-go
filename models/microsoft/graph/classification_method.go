package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "PATTERNMATCH":
            return PATTERNMATCH_CLASSIFICATIONMETHOD, nil
        case "EXACTDATAMATCH":
            return EXACTDATAMATCH_CLASSIFICATIONMETHOD, nil
        case "FINGERPRINT":
            return FINGERPRINT_CLASSIFICATIONMETHOD, nil
        case "MACHINELEARNING":
            return MACHINELEARNING_CLASSIFICATIONMETHOD, nil
    }
    return 0, errors.New("Unknown ClassificationMethod value: " + v)
}
func SerializeClassificationMethod(values []ClassificationMethod) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
