package graph
import (
    "strings"
    "errors"
)
// 
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
    result := CREATE_CLOUDPCAUDITACTIVITYOPERATIONTYPE
    switch strings.ToUpper(v) {
        case "CREATE":
            result = CREATE_CLOUDPCAUDITACTIVITYOPERATIONTYPE
        case "DELETE":
            result = DELETE_CLOUDPCAUDITACTIVITYOPERATIONTYPE
        case "PATCH":
            result = PATCH_CLOUDPCAUDITACTIVITYOPERATIONTYPE
        case "OTHER":
            result = OTHER_CLOUDPCAUDITACTIVITYOPERATIONTYPE
        default:
            return 0, errors.New("Unknown CloudPcAuditActivityOperationType value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcAuditActivityOperationType(values []CloudPcAuditActivityOperationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
