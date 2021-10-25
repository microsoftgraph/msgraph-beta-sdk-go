package graph
import (
    "strings"
    "errors"
)
type CloudPcAuditActivityOperationType int

const (
    CREATE_CLOUDPCAUDITACTIVITYOPERATIONTYPE CloudPcAuditActivityOperationType = iota
    DELETE_CLOUDPCAUDITACTIVITYOPERATIONTYPE
    PATCH_CLOUDPCAUDITACTIVITYOPERATIONTYPE
    OTHER_CLOUDPCAUDITACTIVITYOPERATIONTYPE
)

func (i CloudPcAuditActivityOperationType) String() string {
    return []string{"CREATE", "DELETE", "PATCH", "OTHER"}[i]
}
func ParseCloudPcAuditActivityOperationType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "CREATE":
            return CREATE_CLOUDPCAUDITACTIVITYOPERATIONTYPE, nil
        case "DELETE":
            return DELETE_CLOUDPCAUDITACTIVITYOPERATIONTYPE, nil
        case "PATCH":
            return PATCH_CLOUDPCAUDITACTIVITYOPERATIONTYPE, nil
        case "OTHER":
            return OTHER_CLOUDPCAUDITACTIVITYOPERATIONTYPE, nil
    }
    return 0, errors.New("Unknown CloudPcAuditActivityOperationType value: " + v)
}
func SerializeCloudPcAuditActivityOperationType(values []CloudPcAuditActivityOperationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
