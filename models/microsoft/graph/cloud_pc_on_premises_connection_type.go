package graph
import (
    "strings"
    "errors"
)
// 
type CloudPcOnPremisesConnectionType int

const (
    HYBRIDAZUREADJOIN_CLOUDPCONPREMISESCONNECTIONTYPE CloudPcOnPremisesConnectionType = iota
    AZUREADJOIN_CLOUDPCONPREMISESCONNECTIONTYPE
    UNKNOWNFUTUREVALUE_CLOUDPCONPREMISESCONNECTIONTYPE
)

func (i CloudPcOnPremisesConnectionType) String() string {
    return []string{"HYBRIDAZUREADJOIN", "AZUREADJOIN", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseCloudPcOnPremisesConnectionType(v string) (interface{}, error) {
    result := HYBRIDAZUREADJOIN_CLOUDPCONPREMISESCONNECTIONTYPE
    switch strings.ToUpper(v) {
        case "HYBRIDAZUREADJOIN":
            result = HYBRIDAZUREADJOIN_CLOUDPCONPREMISESCONNECTIONTYPE
        case "AZUREADJOIN":
            result = AZUREADJOIN_CLOUDPCONPREMISESCONNECTIONTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_CLOUDPCONPREMISESCONNECTIONTYPE
        default:
            return 0, errors.New("Unknown CloudPcOnPremisesConnectionType value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcOnPremisesConnectionType(values []CloudPcOnPremisesConnectionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
