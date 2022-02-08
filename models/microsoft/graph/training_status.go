package graph
import (
    "strings"
    "errors"
)
// 
type TrainingStatus int

const (
    UNKNOWN_TRAININGSTATUS TrainingStatus = iota
    ASSIGNED_TRAININGSTATUS
    INPROGRESS_TRAININGSTATUS
    COMPLETED_TRAININGSTATUS
    OVERDUE_TRAININGSTATUS
    UNKNOWNFUTUREVALUE_TRAININGSTATUS
)

func (i TrainingStatus) String() string {
    return []string{"UNKNOWN", "ASSIGNED", "INPROGRESS", "COMPLETED", "OVERDUE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseTrainingStatus(v string) (interface{}, error) {
    result := UNKNOWN_TRAININGSTATUS
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_TRAININGSTATUS
        case "ASSIGNED":
            result = ASSIGNED_TRAININGSTATUS
        case "INPROGRESS":
            result = INPROGRESS_TRAININGSTATUS
        case "COMPLETED":
            result = COMPLETED_TRAININGSTATUS
        case "OVERDUE":
            result = OVERDUE_TRAININGSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_TRAININGSTATUS
        default:
            return 0, errors.New("Unknown TrainingStatus value: " + v)
    }
    return &result, nil
}
func SerializeTrainingStatus(values []TrainingStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
