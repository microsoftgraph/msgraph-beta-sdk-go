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
    NOTCOMPLETED_TRAININGSTATUS
    UNKNOWNFUTUREVALUE_TRAININGSTATUS
)

func (i TrainingStatus) String() string {
    return []string{"UNKNOWN", "ASSIGNED", "INPROGRESS", "COMPLETED", "OVERDUE", "NOTCOMPLETED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseTrainingStatus(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_TRAININGSTATUS, nil
        case "ASSIGNED":
            return ASSIGNED_TRAININGSTATUS, nil
        case "INPROGRESS":
            return INPROGRESS_TRAININGSTATUS, nil
        case "COMPLETED":
            return COMPLETED_TRAININGSTATUS, nil
        case "OVERDUE":
            return OVERDUE_TRAININGSTATUS, nil
        case "NOTCOMPLETED":
            return NOTCOMPLETED_TRAININGSTATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_TRAININGSTATUS, nil
    }
    return 0, errors.New("Unknown TrainingStatus value: " + v)
}
func SerializeTrainingStatus(values []TrainingStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
