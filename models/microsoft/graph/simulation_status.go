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
    RUNNING_SIMULATIONSTATUS
    SCHEDULED_SIMULATIONSTATUS
    SUCCEEDED_SIMULATIONSTATUS
    FAILED_SIMULATIONSTATUS
    CANCELLED_SIMULATIONSTATUS
    EXCLUDED_SIMULATIONSTATUS
    UNKNOWNFUTUREVALUE_SIMULATIONSTATUS
)

func (i SimulationStatus) String() string {
    return []string{"UNKNOWN", "DRAFT", "RUNNING", "SCHEDULED", "SUCCEEDED", "FAILED", "CANCELLED", "EXCLUDED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseSimulationStatus(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_SIMULATIONSTATUS, nil
        case "DRAFT":
            return DRAFT_SIMULATIONSTATUS, nil
        case "RUNNING":
            return RUNNING_SIMULATIONSTATUS, nil
        case "SCHEDULED":
            return SCHEDULED_SIMULATIONSTATUS, nil
        case "SUCCEEDED":
            return SUCCEEDED_SIMULATIONSTATUS, nil
        case "FAILED":
            return FAILED_SIMULATIONSTATUS, nil
        case "CANCELLED":
            return CANCELLED_SIMULATIONSTATUS, nil
        case "EXCLUDED":
            return EXCLUDED_SIMULATIONSTATUS, nil
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
