package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_GROUPPOLICYOPERATIONSTATUS, nil
        case "INPROGRESS":
            return INPROGRESS_GROUPPOLICYOPERATIONSTATUS, nil
        case "SUCCESS":
            return SUCCESS_GROUPPOLICYOPERATIONSTATUS, nil
        case "FAILED":
            return FAILED_GROUPPOLICYOPERATIONSTATUS, nil
    }
    return 0, errors.New("Unknown GroupPolicyOperationStatus value: " + v)
}
func SerializeGroupPolicyOperationStatus(values []GroupPolicyOperationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
