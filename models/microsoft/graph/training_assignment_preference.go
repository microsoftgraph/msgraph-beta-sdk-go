package graph
import (
    "strings"
    "errors"
)
// 
type TrainingAssignmentPreference int

const (
    UNKNOWN_TRAININGASSIGNMENTPREFERENCE TrainingAssignmentPreference = iota
    AUTO_TRAININGASSIGNMENTPREFERENCE
    MANUAL_TRAININGASSIGNMENTPREFERENCE
    UNKNOWNFUTUREVALUE_TRAININGASSIGNMENTPREFERENCE
)

func (i TrainingAssignmentPreference) String() string {
    return []string{"UNKNOWN", "AUTO", "MANUAL", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseTrainingAssignmentPreference(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_TRAININGASSIGNMENTPREFERENCE, nil
        case "AUTO":
            return AUTO_TRAININGASSIGNMENTPREFERENCE, nil
        case "MANUAL":
            return MANUAL_TRAININGASSIGNMENTPREFERENCE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_TRAININGASSIGNMENTPREFERENCE, nil
    }
    return 0, errors.New("Unknown TrainingAssignmentPreference value: " + v)
}
func SerializeTrainingAssignmentPreference(values []TrainingAssignmentPreference) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
