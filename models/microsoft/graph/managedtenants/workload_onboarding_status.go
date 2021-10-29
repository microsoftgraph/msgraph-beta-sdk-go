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
    switch strings.ToUpper(v) {
        case "NOTONBOARDED":
            return NOTONBOARDED_WORKLOADONBOARDINGSTATUS, nil
        case "ONBOARDED":
            return ONBOARDED_WORKLOADONBOARDINGSTATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_WORKLOADONBOARDINGSTATUS, nil
    }
    return 0, errors.New("Unknown WorkloadOnboardingStatus value: " + v)
}
func SerializeWorkloadOnboardingStatus(values []WorkloadOnboardingStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
