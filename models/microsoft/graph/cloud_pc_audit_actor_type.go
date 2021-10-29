package graph
import (
    "strings"
    "errors"
)
// 
type CloudPcAuditActorType int

const (
    ITPRO_CLOUDPCAUDITACTORTYPE CloudPcAuditActorType = iota
    APPLICATION_CLOUDPCAUDITACTORTYPE
    PARTNER_CLOUDPCAUDITACTORTYPE
    UNKNOWN_CLOUDPCAUDITACTORTYPE
)

func (i CloudPcAuditActorType) String() string {
    return []string{"ITPRO", "APPLICATION", "PARTNER", "UNKNOWN"}[i]
}
func ParseCloudPcAuditActorType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "ITPRO":
            return ITPRO_CLOUDPCAUDITACTORTYPE, nil
        case "APPLICATION":
            return APPLICATION_CLOUDPCAUDITACTORTYPE, nil
        case "PARTNER":
            return PARTNER_CLOUDPCAUDITACTORTYPE, nil
        case "UNKNOWN":
            return UNKNOWN_CLOUDPCAUDITACTORTYPE, nil
    }
    return 0, errors.New("Unknown CloudPcAuditActorType value: " + v)
}
func SerializeCloudPcAuditActorType(values []CloudPcAuditActorType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
