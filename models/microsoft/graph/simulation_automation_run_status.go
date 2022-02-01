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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_SIMULATIONAUTOMATIONRUNSTATUS, nil
        case "RUNNING":
            return RUNNING_SIMULATIONAUTOMATIONRUNSTATUS, nil
        case "SUCCEEDED":
            return SUCCEEDED_SIMULATIONAUTOMATIONRUNSTATUS, nil
        case "FAILED":
            return FAILED_SIMULATIONAUTOMATIONRUNSTATUS, nil
        case "SKIPPED":
            return SKIPPED_SIMULATIONAUTOMATIONRUNSTATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_SIMULATIONAUTOMATIONRUNSTATUS, nil
    }
    return 0, errors.New("Unknown SimulationAutomationRunStatus value: " + v)
}
func SerializeSimulationAutomationRunStatus(values []SimulationAutomationRunStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
