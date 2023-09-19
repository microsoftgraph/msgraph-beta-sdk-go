package managedtenants
import (
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
    return []string{"notOnboarded", "onboarded", "unknownFutureValue"}[i]
}
func ParseWorkloadOnboardingStatus(v string) (any, error) {
    result := NOTONBOARDED_WORKLOADONBOARDINGSTATUS
    switch v {
        case "notOnboarded":
            result = NOTONBOARDED_WORKLOADONBOARDINGSTATUS
        case "onboarded":
            result = ONBOARDED_WORKLOADONBOARDINGSTATUS
        case "unknownFutureValue":
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
func (i WorkloadOnboardingStatus) isMultiValue() bool {
    return false
}
