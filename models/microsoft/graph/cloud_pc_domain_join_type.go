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
    switch strings.ToUpper(v) {
        case "AZUREADJOIN":
            return AZUREADJOIN_CLOUDPCDOMAINJOINTYPE, nil
        case "HYBRIDAZUREADJOIN":
            return HYBRIDAZUREADJOIN_CLOUDPCDOMAINJOINTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_CLOUDPCDOMAINJOINTYPE, nil
    }
    return 0, errors.New("Unknown CloudPcDomainJoinType value: " + v)
}
func SerializeCloudPcDomainJoinType(values []CloudPcDomainJoinType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
