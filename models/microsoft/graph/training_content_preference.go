package graph
import (
    "strings"
    "errors"
)
// 
type TrainingContentPreference int

const (
    UNKNOWN_TRAININGCONTENTPREFERENCE TrainingContentPreference = iota
    MICROSOFT_TRAININGCONTENTPREFERENCE
    CUSTOM_TRAININGCONTENTPREFERENCE
    NOTRAINING_TRAININGCONTENTPREFERENCE
    UNKNOWNFUTUREVALUE_TRAININGCONTENTPREFERENCE
)

func (i TrainingContentPreference) String() string {
    return []string{"UNKNOWN", "MICROSOFT", "CUSTOM", "NOTRAINING", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseTrainingContentPreference(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_TRAININGCONTENTPREFERENCE, nil
        case "MICROSOFT":
            return MICROSOFT_TRAININGCONTENTPREFERENCE, nil
        case "CUSTOM":
            return CUSTOM_TRAININGCONTENTPREFERENCE, nil
        case "NOTRAINING":
            return NOTRAINING_TRAININGCONTENTPREFERENCE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_TRAININGCONTENTPREFERENCE, nil
    }
    return 0, errors.New("Unknown TrainingContentPreference value: " + v)
}
func SerializeTrainingContentPreference(values []TrainingContentPreference) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
