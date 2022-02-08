package graph
import (
    "strings"
    "errors"
)
// 
type KerberosSignOnMappingAttributeType int

const (
    USERPRINCIPALNAME_KERBEROSSIGNONMAPPINGATTRIBUTETYPE KerberosSignOnMappingAttributeType = iota
    ONPREMISESUSERPRINCIPALNAME_KERBEROSSIGNONMAPPINGATTRIBUTETYPE
    USERPRINCIPALUSERNAME_KERBEROSSIGNONMAPPINGATTRIBUTETYPE
    ONPREMISESUSERPRINCIPALUSERNAME_KERBEROSSIGNONMAPPINGATTRIBUTETYPE
    ONPREMISESSAMACCOUNTNAME_KERBEROSSIGNONMAPPINGATTRIBUTETYPE
)

func (i KerberosSignOnMappingAttributeType) String() string {
    return []string{"USERPRINCIPALNAME", "ONPREMISESUSERPRINCIPALNAME", "USERPRINCIPALUSERNAME", "ONPREMISESUSERPRINCIPALUSERNAME", "ONPREMISESSAMACCOUNTNAME"}[i]
}
func ParseKerberosSignOnMappingAttributeType(v string) (interface{}, error) {
    result := USERPRINCIPALNAME_KERBEROSSIGNONMAPPINGATTRIBUTETYPE
    switch strings.ToUpper(v) {
        case "USERPRINCIPALNAME":
            result = USERPRINCIPALNAME_KERBEROSSIGNONMAPPINGATTRIBUTETYPE
        case "ONPREMISESUSERPRINCIPALNAME":
            result = ONPREMISESUSERPRINCIPALNAME_KERBEROSSIGNONMAPPINGATTRIBUTETYPE
        case "USERPRINCIPALUSERNAME":
            result = USERPRINCIPALUSERNAME_KERBEROSSIGNONMAPPINGATTRIBUTETYPE
        case "ONPREMISESUSERPRINCIPALUSERNAME":
            result = ONPREMISESUSERPRINCIPALUSERNAME_KERBEROSSIGNONMAPPINGATTRIBUTETYPE
        case "ONPREMISESSAMACCOUNTNAME":
            result = ONPREMISESSAMACCOUNTNAME_KERBEROSSIGNONMAPPINGATTRIBUTETYPE
        default:
            return 0, errors.New("Unknown KerberosSignOnMappingAttributeType value: " + v)
    }
    return &result, nil
}
func SerializeKerberosSignOnMappingAttributeType(values []KerberosSignOnMappingAttributeType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
