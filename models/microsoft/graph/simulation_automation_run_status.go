package graph
import (
    "strings"
    "errors"
)
// 
type SimulationAutomationRunStatus int

const (
    UNKNOWN_SIMULATIONAUTOMATIONRUNSTATUS SimulationAutomationRunStatus = iota
    RUNNING_SIMULATIONAUTOMATIONRUNSTATUS
    SUCCEEDED_SIMULATIONAUTOMATIONRUNSTATUS
    FAILED_SIMULATIONAUTOMATIONRUNSTATUS
    SKIPPED_SIMULATIONAUTOMATIONRUNSTATUS
    UNKNOWNFUTUREVALUE_SIMULATIONAUTOMATIONRUNSTATUS
)

func (i SimulationAutomationRunStatus) String() string {
    return []string{"UNKNOWN", "RUNNING", "SUCCEEDED", "FAILED", "SKIPPED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseSimulationAutomationRunStatus(v string) (interface{}, error) {
    result := UNKNOWN_SIMULATIONAUTOMATIONRUNSTATUS
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_SIMULATIONAUTOMATIONRUNSTATUS
        case "RUNNING":
            result = RUNNING_SIMULATIONAUTOMATIONRUNSTATUS
        case "SUCCEEDED":
            result = SUCCEEDED_SIMULATIONAUTOMATIONRUNSTATUS
        case "FAILED":
            result = FAILED_SIMULATIONAUTOMATIONRUNSTATUS
        case "SKIPPED":
            result = SKIPPED_SIMULATIONAUTOMATIONRUNSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_SIMULATIONAUTOMATIONRUNSTATUS
        default:
            return 0, errors.New("Unknown SimulationAutomationRunStatus value: " + v)
    }
    return &result, nil
}
func SerializeSimulationAutomationRunStatus(values []SimulationAutomationRunStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
