// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models
type ClassificationMethod int

const (
    PATTERNMATCH_CLASSIFICATIONMETHOD ClassificationMethod = iota
    EXACTDATAMATCH_CLASSIFICATIONMETHOD
    FINGERPRINT_CLASSIFICATIONMETHOD
    MACHINELEARNING_CLASSIFICATIONMETHOD
)

func (i ClassificationMethod) String() string {
    return []string{"patternMatch", "exactDataMatch", "fingerprint", "machineLearning"}[i]
}
func ParseClassificationMethod(v string) (any, error) {
    result := PATTERNMATCH_CLASSIFICATIONMETHOD
    switch v {
        case "patternMatch":
            result = PATTERNMATCH_CLASSIFICATIONMETHOD
        case "exactDataMatch":
            result = EXACTDATAMATCH_CLASSIFICATIONMETHOD
        case "fingerprint":
            result = FINGERPRINT_CLASSIFICATIONMETHOD
        case "machineLearning":
            result = MACHINELEARNING_CLASSIFICATIONMETHOD
        default:
            return nil, nil
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
func (i ClassificationMethod) isMultiValue() bool {
    return false
}
