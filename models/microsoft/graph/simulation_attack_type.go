package graph
import (
    "strings"
    "errors"
)
// 
type SimulationAttackType int

const (
    UNKNOWN_SIMULATIONATTACKTYPE SimulationAttackType = iota
    SOCIAL_SIMULATIONATTACKTYPE
    CLOUD_SIMULATIONATTACKTYPE
    ENDPOINT_SIMULATIONATTACKTYPE
    UNKNOWNFUTUREVALUE_SIMULATIONATTACKTYPE
)

func (i SimulationAttackType) String() string {
    return []string{"UNKNOWN", "SOCIAL", "CLOUD", "ENDPOINT", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseSimulationAttackType(v string) (interface{}, error) {
    result := UNKNOWN_SIMULATIONATTACKTYPE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_SIMULATIONATTACKTYPE
        case "SOCIAL":
            result = SOCIAL_SIMULATIONATTACKTYPE
        case "CLOUD":
            result = CLOUD_SIMULATIONATTACKTYPE
        case "ENDPOINT":
            result = ENDPOINT_SIMULATIONATTACKTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_SIMULATIONATTACKTYPE
        default:
            return 0, errors.New("Unknown SimulationAttackType value: " + v)
    }
    return &result, nil
}
func SerializeSimulationAttackType(values []SimulationAttackType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
