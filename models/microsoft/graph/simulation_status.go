package graph
import (
    "strings"
    "errors"
)
// 
type SimulationStatus int

const (
    UNKNOWN_SIMULATIONSTATUS SimulationStatus = iota
    DRAFT_SIMULATIONSTATUS
    INPROGRESS_SIMULATIONSTATUS
    SCHEDULED_SIMULATIONSTATUS
    COMPLETED_SIMULATIONSTATUS
    PARTIALLYCOMPLETED_SIMULATIONSTATUS
    FAILED_SIMULATIONSTATUS
    CANCELLED_SIMULATIONSTATUS
    EXCLUDED_SIMULATIONSTATUS
    DELETED_SIMULATIONSTATUS
    INCLUDED_SIMULATIONSTATUS
    UNKNOWNFUTUREVALUE_SIMULATIONSTATUS
)

func (i SimulationStatus) String() string {
    return []string{"UNKNOWN", "DRAFT", "INPROGRESS", "SCHEDULED", "COMPLETED", "PARTIALLYCOMPLETED", "FAILED", "CANCELLED", "EXCLUDED", "DELETED", "INCLUDED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseSimulationStatus(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_SIMULATIONSTATUS, nil
        case "DRAFT":
            return DRAFT_SIMULATIONSTATUS, nil
        case "INPROGRESS":
            return INPROGRESS_SIMULATIONSTATUS, nil
        case "SCHEDULED":
            return SCHEDULED_SIMULATIONSTATUS, nil
        case "COMPLETED":
            return COMPLETED_SIMULATIONSTATUS, nil
        case "PARTIALLYCOMPLETED":
            return PARTIALLYCOMPLETED_SIMULATIONSTATUS, nil
        case "FAILED":
            return FAILED_SIMULATIONSTATUS, nil
        case "CANCELLED":
            return CANCELLED_SIMULATIONSTATUS, nil
        case "EXCLUDED":
            return EXCLUDED_SIMULATIONSTATUS, nil
        case "DELETED":
            return DELETED_SIMULATIONSTATUS, nil
        case "INCLUDED":
            return INCLUDED_SIMULATIONSTATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_SIMULATIONSTATUS, nil
    }
    return 0, errors.New("Unknown SimulationStatus value: " + v)
}
func SerializeSimulationStatus(values []SimulationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
