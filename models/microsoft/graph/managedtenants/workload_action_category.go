package managedtenants
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "AUTOMATED":
            return AUTOMATED_WORKLOADACTIONCATEGORY, nil
        case "MANUAL":
            return MANUAL_WORKLOADACTIONCATEGORY, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_WORKLOADACTIONCATEGORY, nil
    }
    return 0, errors.New("Unknown WorkloadActionCategory value: " + v)
}
func SerializeWorkloadActionCategory(values []WorkloadActionCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
