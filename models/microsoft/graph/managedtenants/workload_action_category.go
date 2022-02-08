package managedtenants
import (
    "strings"
    "errors"
)
// 
type WorkloadActionCategory int

const (
    AUTOMATED_WORKLOADACTIONCATEGORY WorkloadActionCategory = iota
    MANUAL_WORKLOADACTIONCATEGORY
    UNKNOWNFUTUREVALUE_WORKLOADACTIONCATEGORY
)

func (i WorkloadActionCategory) String() string {
    return []string{"AUTOMATED", "MANUAL", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseWorkloadActionCategory(v string) (interface{}, error) {
    result := AUTOMATED_WORKLOADACTIONCATEGORY
    switch strings.ToUpper(v) {
        case "AUTOMATED":
            result = AUTOMATED_WORKLOADACTIONCATEGORY
        case "MANUAL":
            result = MANUAL_WORKLOADACTIONCATEGORY
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_WORKLOADACTIONCATEGORY
        default:
            return 0, errors.New("Unknown WorkloadActionCategory value: " + v)
    }
    return &result, nil
}
func SerializeWorkloadActionCategory(values []WorkloadActionCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
