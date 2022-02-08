package graph
import (
    "strings"
    "errors"
)
// 
type SimulationAttackTechnique int

const (
    UNKNOWN_SIMULATIONATTACKTECHNIQUE SimulationAttackTechnique = iota
    CREDENTIALHARVESTING_SIMULATIONATTACKTECHNIQUE
    ATTACHMENTMALWARE_SIMULATIONATTACKTECHNIQUE
    DRIVEBYURL_SIMULATIONATTACKTECHNIQUE
    LINKINATTACHMENT_SIMULATIONATTACKTECHNIQUE
    LINKTOMALWAREFILE_SIMULATIONATTACKTECHNIQUE
    UNKNOWNFUTUREVALUE_SIMULATIONATTACKTECHNIQUE
)

func (i SimulationAttackTechnique) String() string {
    return []string{"UNKNOWN", "CREDENTIALHARVESTING", "ATTACHMENTMALWARE", "DRIVEBYURL", "LINKINATTACHMENT", "LINKTOMALWAREFILE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseSimulationAttackTechnique(v string) (interface{}, error) {
    result := UNKNOWN_SIMULATIONATTACKTECHNIQUE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_SIMULATIONATTACKTECHNIQUE
        case "CREDENTIALHARVESTING":
            result = CREDENTIALHARVESTING_SIMULATIONATTACKTECHNIQUE
        case "ATTACHMENTMALWARE":
            result = ATTACHMENTMALWARE_SIMULATIONATTACKTECHNIQUE
        case "DRIVEBYURL":
            result = DRIVEBYURL_SIMULATIONATTACKTECHNIQUE
        case "LINKINATTACHMENT":
            result = LINKINATTACHMENT_SIMULATIONATTACKTECHNIQUE
        case "LINKTOMALWAREFILE":
            result = LINKTOMALWAREFILE_SIMULATIONATTACKTECHNIQUE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_SIMULATIONATTACKTECHNIQUE
        default:
            return 0, errors.New("Unknown SimulationAttackTechnique value: " + v)
    }
    return &result, nil
}
func SerializeSimulationAttackTechnique(values []SimulationAttackTechnique) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
