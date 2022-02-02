package graph
import (
    "strings"
    "errors"
)
// 
type SimulationAutomationStatus int

const (
    UNKNOWN_SIMULATIONAUTOMATIONSTATUS SimulationAutomationStatus = iota
    DRAFT_SIMULATIONAUTOMATIONSTATUS
    NOTRUNNING_SIMULATIONAUTOMATIONSTATUS
    RUNNING_SIMULATIONAUTOMATIONSTATUS
    COMPLETED_SIMULATIONAUTOMATIONSTATUS
    UNKNOWNFUTUREVALUE_SIMULATIONAUTOMATIONSTATUS
)

func (i SimulationAutomationStatus) String() string {
    return []string{"UNKNOWN", "DRAFT", "NOTRUNNING", "RUNNING", "COMPLETED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseSimulationAutomationStatus(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_SIMULATIONAUTOMATIONSTATUS, nil
        case "DRAFT":
            return DRAFT_SIMULATIONAUTOMATIONSTATUS, nil
        case "NOTRUNNING":
            return NOTRUNNING_SIMULATIONAUTOMATIONSTATUS, nil
        case "RUNNING":
            return RUNNING_SIMULATIONAUTOMATIONSTATUS, nil
        case "COMPLETED":
            return COMPLETED_SIMULATIONAUTOMATIONSTATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_SIMULATIONAUTOMATIONSTATUS, nil
    }
    return 0, errors.New("Unknown SimulationAutomationStatus value: " + v)
}
func SerializeSimulationAutomationStatus(values []SimulationAutomationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
