package graph
import (
    "strings"
    "errors"
)
type OnboardingStatus int

const (
    UNKNOWN_ONBOARDINGSTATUS OnboardingStatus = iota
    INPROGRESS_ONBOARDINGSTATUS
    ONBOARDED_ONBOARDINGSTATUS
    FAILED_ONBOARDINGSTATUS
)

func (i OnboardingStatus) String() string {
    return []string{"UNKNOWN", "INPROGRESS", "ONBOARDED", "FAILED"}[i]
}
func ParseOnboardingStatus(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_ONBOARDINGSTATUS, nil
        case "INPROGRESS":
            return INPROGRESS_ONBOARDINGSTATUS, nil
        case "ONBOARDED":
            return ONBOARDED_ONBOARDINGSTATUS, nil
        case "FAILED":
            return FAILED_ONBOARDINGSTATUS, nil
    }
    return 0, errors.New("Unknown OnboardingStatus value: " + v)
}
func SerializeOnboardingStatus(values []OnboardingStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
