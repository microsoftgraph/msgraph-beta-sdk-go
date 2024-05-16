package managedtenants
type WorkloadActionCategory int

const (
    AUTOMATED_WORKLOADACTIONCATEGORY WorkloadActionCategory = iota
    MANUAL_WORKLOADACTIONCATEGORY
    UNKNOWNFUTUREVALUE_WORKLOADACTIONCATEGORY
)

func (i WorkloadActionCategory) String() string {
    return []string{"automated", "manual", "unknownFutureValue"}[i]
}
func ParseWorkloadActionCategory(v string) (any, error) {
    result := AUTOMATED_WORKLOADACTIONCATEGORY
    switch v {
        case "automated":
            result = AUTOMATED_WORKLOADACTIONCATEGORY
        case "manual":
            result = MANUAL_WORKLOADACTIONCATEGORY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_WORKLOADACTIONCATEGORY
        default:
            return nil, nil
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
func (i WorkloadActionCategory) isMultiValue() bool {
    return false
}
