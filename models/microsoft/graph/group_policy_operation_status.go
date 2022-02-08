package graph
import (
    "strings"
    "errors"
)
// 
type GroupPolicyOperationStatus int

const (
    UNKNOWN_GROUPPOLICYOPERATIONSTATUS GroupPolicyOperationStatus = iota
    INPROGRESS_GROUPPOLICYOPERATIONSTATUS
    SUCCESS_GROUPPOLICYOPERATIONSTATUS
    FAILED_GROUPPOLICYOPERATIONSTATUS
)

func (i GroupPolicyOperationStatus) String() string {
    return []string{"UNKNOWN", "INPROGRESS", "SUCCESS", "FAILED"}[i]
}
func ParseGroupPolicyOperationStatus(v string) (interface{}, error) {
    result := UNKNOWN_GROUPPOLICYOPERATIONSTATUS
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_GROUPPOLICYOPERATIONSTATUS
        case "INPROGRESS":
            result = INPROGRESS_GROUPPOLICYOPERATIONSTATUS
        case "SUCCESS":
            result = SUCCESS_GROUPPOLICYOPERATIONSTATUS
        case "FAILED":
            result = FAILED_GROUPPOLICYOPERATIONSTATUS
        default:
            return 0, errors.New("Unknown GroupPolicyOperationStatus value: " + v)
    }
    return &result, nil
}
func SerializeGroupPolicyOperationStatus(values []GroupPolicyOperationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
