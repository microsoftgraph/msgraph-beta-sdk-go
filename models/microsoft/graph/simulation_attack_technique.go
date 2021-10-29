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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_SIMULATIONATTACKTECHNIQUE, nil
        case "CREDENTIALHARVESTING":
            return CREDENTIALHARVESTING_SIMULATIONATTACKTECHNIQUE, nil
        case "ATTACHMENTMALWARE":
            return ATTACHMENTMALWARE_SIMULATIONATTACKTECHNIQUE, nil
        case "DRIVEBYURL":
            return DRIVEBYURL_SIMULATIONATTACKTECHNIQUE, nil
        case "LINKINATTACHMENT":
            return LINKINATTACHMENT_SIMULATIONATTACKTECHNIQUE, nil
        case "LINKTOMALWAREFILE":
            return LINKTOMALWAREFILE_SIMULATIONATTACKTECHNIQUE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_SIMULATIONATTACKTECHNIQUE, nil
    }
    return 0, errors.New("Unknown SimulationAttackTechnique value: " + v)
}
func SerializeSimulationAttackTechnique(values []SimulationAttackTechnique) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
