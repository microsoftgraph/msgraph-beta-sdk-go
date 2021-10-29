package graph
import (
    "strings"
    "errors"
)
// 
type SimulationMode int

const (
    REAL_SIMULATIONMODE SimulationMode = iota
    PREVIEW_SIMULATIONMODE
    UNKNOWNFUTUREVALUE_SIMULATIONMODE
)

func (i SimulationMode) String() string {
    return []string{"REAL", "PREVIEW", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseSimulationMode(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "REAL":
            return REAL_SIMULATIONMODE, nil
        case "PREVIEW":
            return PREVIEW_SIMULATIONMODE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_SIMULATIONMODE, nil
    }
    return 0, errors.New("Unknown SimulationMode value: " + v)
}
func SerializeSimulationMode(values []SimulationMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
