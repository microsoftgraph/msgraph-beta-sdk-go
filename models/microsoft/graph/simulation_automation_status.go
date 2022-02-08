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
    result := UNKNOWN_SIMULATIONAUTOMATIONSTATUS
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_SIMULATIONAUTOMATIONSTATUS
        case "DRAFT":
            result = DRAFT_SIMULATIONAUTOMATIONSTATUS
        case "NOTRUNNING":
            result = NOTRUNNING_SIMULATIONAUTOMATIONSTATUS
        case "RUNNING":
            result = RUNNING_SIMULATIONAUTOMATIONSTATUS
        case "COMPLETED":
            result = COMPLETED_SIMULATIONAUTOMATIONSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_SIMULATIONAUTOMATIONSTATUS
        default:
            return 0, errors.New("Unknown SimulationAutomationStatus value: " + v)
    }
    return &result, nil
}
func SerializeSimulationAutomationStatus(values []SimulationAutomationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
