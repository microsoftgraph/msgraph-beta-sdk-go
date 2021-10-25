package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "SUCCESS":
            return SUCCESS_CLOUDPCAUDITACTIVITYRESULT, nil
        case "CLIENTERROR":
            return CLIENTERROR_CLOUDPCAUDITACTIVITYRESULT, nil
        case "FAILURE":
            return FAILURE_CLOUDPCAUDITACTIVITYRESULT, nil
        case "TIMEOUT":
            return TIMEOUT_CLOUDPCAUDITACTIVITYRESULT, nil
        case "OTHER":
            return OTHER_CLOUDPCAUDITACTIVITYRESULT, nil
    }
    return 0, errors.New("Unknown CloudPcAuditActivityResult value: " + v)
}
func SerializeCloudPcAuditActivityResult(values []CloudPcAuditActivityResult) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
