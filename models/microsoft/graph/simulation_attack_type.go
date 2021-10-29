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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_SIMULATIONATTACKTYPE, nil
        case "SOCIAL":
            return SOCIAL_SIMULATIONATTACKTYPE, nil
        case "CLOUD":
            return CLOUD_SIMULATIONATTACKTYPE, nil
        case "ENDPOINT":
            return ENDPOINT_SIMULATIONATTACKTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_SIMULATIONATTACKTYPE, nil
    }
    return 0, errors.New("Unknown SimulationAttackType value: " + v)
}
func SerializeSimulationAttackType(values []SimulationAttackType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
