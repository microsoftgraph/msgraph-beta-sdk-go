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
    result := ITPRO_CLOUDPCAUDITACTORTYPE
    switch strings.ToUpper(v) {
        case "ITPRO":
            result = ITPRO_CLOUDPCAUDITACTORTYPE
        case "APPLICATION":
            result = APPLICATION_CLOUDPCAUDITACTORTYPE
        case "PARTNER":
            result = PARTNER_CLOUDPCAUDITACTORTYPE
        case "UNKNOWN":
            result = UNKNOWN_CLOUDPCAUDITACTORTYPE
        default:
            return 0, errors.New("Unknown CloudPcAuditActorType value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcAuditActorType(values []CloudPcAuditActorType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
