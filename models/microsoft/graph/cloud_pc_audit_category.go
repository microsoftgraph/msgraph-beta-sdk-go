package graph
import (
    "strings"
    "errors"
)
type CloudPcAuditCategory int

const (
    CLOUDPC_CLOUDPCAUDITCATEGORY CloudPcAuditCategory = iota
    OTHER_CLOUDPCAUDITCATEGORY
)

func (i CloudPcAuditCategory) String() string {
    return []string{"CLOUDPC", "OTHER"}[i]
}
func ParseCloudPcAuditCategory(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "CLOUDPC":
            return CLOUDPC_CLOUDPCAUDITCATEGORY, nil
        case "OTHER":
            return OTHER_CLOUDPCAUDITCATEGORY, nil
    }
    return 0, errors.New("Unknown CloudPcAuditCategory value: " + v)
}
func SerializeCloudPcAuditCategory(values []CloudPcAuditCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
