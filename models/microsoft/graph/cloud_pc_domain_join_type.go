package graph
import (
    "strings"
    "errors"
)
// 
type CloudPcDomainJoinType int

const (
    AZUREADJOIN_CLOUDPCDOMAINJOINTYPE CloudPcDomainJoinType = iota
    HYBRIDAZUREADJOIN_CLOUDPCDOMAINJOINTYPE
    UNKNOWNFUTUREVALUE_CLOUDPCDOMAINJOINTYPE
)

func (i CloudPcDomainJoinType) String() string {
    return []string{"AZUREADJOIN", "HYBRIDAZUREADJOIN", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseCloudPcDomainJoinType(v string) (interface{}, error) {
    result := AZUREADJOIN_CLOUDPCDOMAINJOINTYPE
    switch strings.ToUpper(v) {
        case "AZUREADJOIN":
            result = AZUREADJOIN_CLOUDPCDOMAINJOINTYPE
        case "HYBRIDAZUREADJOIN":
            result = HYBRIDAZUREADJOIN_CLOUDPCDOMAINJOINTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_CLOUDPCDOMAINJOINTYPE
        default:
            return 0, errors.New("Unknown CloudPcDomainJoinType value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcDomainJoinType(values []CloudPcDomainJoinType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
