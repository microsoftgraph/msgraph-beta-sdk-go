package graph
import (
    "strings"
    "errors"
)
// 
type CloudPcAuditActivityResult int

const (
    SUCCESS_CLOUDPCAUDITACTIVITYRESULT CloudPcAuditActivityResult = iota
    CLIENTERROR_CLOUDPCAUDITACTIVITYRESULT
    FAILURE_CLOUDPCAUDITACTIVITYRESULT
    TIMEOUT_CLOUDPCAUDITACTIVITYRESULT
    OTHER_CLOUDPCAUDITACTIVITYRESULT
)

func (i CloudPcAuditActivityResult) String() string {
    return []string{"SUCCESS", "CLIENTERROR", "FAILURE", "TIMEOUT", "OTHER"}[i]
}
func ParseCloudPcAuditActivityResult(v string) (interface{}, error) {
    result := SUCCESS_CLOUDPCAUDITACTIVITYRESULT
    switch strings.ToUpper(v) {
        case "SUCCESS":
            result = SUCCESS_CLOUDPCAUDITACTIVITYRESULT
        case "CLIENTERROR":
            result = CLIENTERROR_CLOUDPCAUDITACTIVITYRESULT
        case "FAILURE":
            result = FAILURE_CLOUDPCAUDITACTIVITYRESULT
        case "TIMEOUT":
            result = TIMEOUT_CLOUDPCAUDITACTIVITYRESULT
        case "OTHER":
            result = OTHER_CLOUDPCAUDITACTIVITYRESULT
        default:
            return 0, errors.New("Unknown CloudPcAuditActivityResult value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcAuditActivityResult(values []CloudPcAuditActivityResult) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
