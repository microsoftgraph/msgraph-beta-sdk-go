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
    result := UNKNOWN_SIMULATIONSTATUS
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_SIMULATIONSTATUS
        case "DRAFT":
            result = DRAFT_SIMULATIONSTATUS
        case "RUNNING":
            result = RUNNING_SIMULATIONSTATUS
        case "SCHEDULED":
            result = SCHEDULED_SIMULATIONSTATUS
        case "SUCCEEDED":
            result = SUCCEEDED_SIMULATIONSTATUS
        case "FAILED":
            result = FAILED_SIMULATIONSTATUS
        case "CANCELLED":
            result = CANCELLED_SIMULATIONSTATUS
        case "EXCLUDED":
            result = EXCLUDED_SIMULATIONSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_SIMULATIONSTATUS
        default:
            return 0, errors.New("Unknown SimulationStatus value: " + v)
    }
    return &result, nil
}
func SerializeSimulationStatus(values []SimulationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
