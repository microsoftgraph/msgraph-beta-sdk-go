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
    switch strings.ToUpper(v) {
        case "HYBRIDAZUREADJOIN":
            return HYBRIDAZUREADJOIN_CLOUDPCONPREMISESCONNECTIONTYPE, nil
        case "AZUREADJOIN":
            return AZUREADJOIN_CLOUDPCONPREMISESCONNECTIONTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_CLOUDPCONPREMISESCONNECTIONTYPE, nil
    }
    return 0, errors.New("Unknown CloudPcOnPremisesConnectionType value: " + v)
}
func SerializeCloudPcOnPremisesConnectionType(values []CloudPcOnPremisesConnectionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
