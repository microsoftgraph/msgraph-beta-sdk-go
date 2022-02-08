package managedtenants
import (
    "strings"
    "errors"
)
// 
type WorkloadOnboardingStatus int

const (
    NOTONBOARDED_WORKLOADONBOARDINGSTATUS WorkloadOnboardingStatus = iota
    ONBOARDED_WORKLOADONBOARDINGSTATUS
    UNKNOWNFUTUREVALUE_WORKLOADONBOARDINGSTATUS
)

func (i WorkloadOnboardingStatus) String() string {
    return []string{"NOTONBOARDED", "ONBOARDED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseWorkloadOnboardingStatus(v string) (interface{}, error) {
    result := NOTONBOARDED_WORKLOADONBOARDINGSTATUS
    switch strings.ToUpper(v) {
        case "NOTONBOARDED":
            result = NOTONBOARDED_WORKLOADONBOARDINGSTATUS
        case "ONBOARDED":
            result = ONBOARDED_WORKLOADONBOARDINGSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_WORKLOADONBOARDINGSTATUS
        default:
            return 0, errors.New("Unknown WorkloadOnboardingStatus value: " + v)
    }
    return &result, nil
}
func SerializeWorkloadOnboardingStatus(values []WorkloadOnboardingStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
