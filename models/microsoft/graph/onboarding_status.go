package graph
import (
    "strings"
    "errors"
)
// 
type OnboardingStatus int

const (
    UNKNOWN_ONBOARDINGSTATUS OnboardingStatus = iota
    INPROGRESS_ONBOARDINGSTATUS
    ONBOARDED_ONBOARDINGSTATUS
    FAILED_ONBOARDINGSTATUS
    OFFBOARDING_ONBOARDINGSTATUS
    UNKNOWNFUTUREVALUE_ONBOARDINGSTATUS
)

func (i OnboardingStatus) String() string {
    return []string{"UNKNOWN", "INPROGRESS", "ONBOARDED", "FAILED", "OFFBOARDING", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseOnboardingStatus(v string) (interface{}, error) {
    result := UNKNOWN_ONBOARDINGSTATUS
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_ONBOARDINGSTATUS
        case "INPROGRESS":
            result = INPROGRESS_ONBOARDINGSTATUS
        case "ONBOARDED":
            result = ONBOARDED_ONBOARDINGSTATUS
        case "FAILED":
            result = FAILED_ONBOARDINGSTATUS
        case "OFFBOARDING":
            result = OFFBOARDING_ONBOARDINGSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ONBOARDINGSTATUS
        default:
            return 0, errors.New("Unknown OnboardingStatus value: " + v)
    }
    return &result, nil
}
func SerializeOnboardingStatus(values []OnboardingStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
