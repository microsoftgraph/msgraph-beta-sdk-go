package graph
import (
    "strings"
    "errors"
)
// 
type CloudPcAuditCategory int

const (
    CLOUDPC_CLOUDPCAUDITCATEGORY CloudPcAuditCategory = iota
    OTHER_CLOUDPCAUDITCATEGORY
)

func (i CloudPcAuditCategory) String() string {
    return []string{"CLOUDPC", "OTHER"}[i]
}
func ParseCloudPcAuditCategory(v string) (interface{}, error) {
    result := CLOUDPC_CLOUDPCAUDITCATEGORY
    switch strings.ToUpper(v) {
        case "CLOUDPC":
            result = CLOUDPC_CLOUDPCAUDITCATEGORY
        case "OTHER":
            result = OTHER_CLOUDPCAUDITCATEGORY
        default:
            return 0, errors.New("Unknown CloudPcAuditCategory value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcAuditCategory(values []CloudPcAuditCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
