package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "USERPRINCIPALNAME":
            return USERPRINCIPALNAME_KERBEROSSIGNONMAPPINGATTRIBUTETYPE, nil
        case "ONPREMISESUSERPRINCIPALNAME":
            return ONPREMISESUSERPRINCIPALNAME_KERBEROSSIGNONMAPPINGATTRIBUTETYPE, nil
        case "USERPRINCIPALUSERNAME":
            return USERPRINCIPALUSERNAME_KERBEROSSIGNONMAPPINGATTRIBUTETYPE, nil
        case "ONPREMISESUSERPRINCIPALUSERNAME":
            return ONPREMISESUSERPRINCIPALUSERNAME_KERBEROSSIGNONMAPPINGATTRIBUTETYPE, nil
        case "ONPREMISESSAMACCOUNTNAME":
            return ONPREMISESSAMACCOUNTNAME_KERBEROSSIGNONMAPPINGATTRIBUTETYPE, nil
    }
    return 0, errors.New("Unknown KerberosSignOnMappingAttributeType value: " + v)
}
func SerializeKerberosSignOnMappingAttributeType(values []KerberosSignOnMappingAttributeType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
