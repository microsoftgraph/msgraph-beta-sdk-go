package graph
import (
    "strings"
    "errors"
)
// 
type ChromeOSOnboardingStatus int

const (
    UNKNOWN_CHROMEOSONBOARDINGSTATUS ChromeOSOnboardingStatus = iota
    INPROGRESS_CHROMEOSONBOARDINGSTATUS
    ONBOARDED_CHROMEOSONBOARDINGSTATUS
    FAILED_CHROMEOSONBOARDINGSTATUS
    OFFBOARDING_CHROMEOSONBOARDINGSTATUS
    UNKNOWNFUTUREVALUE_CHROMEOSONBOARDINGSTATUS
)

func (i ChromeOSOnboardingStatus) String() string {
    return []string{"UNKNOWN", "INPROGRESS", "ONBOARDED", "FAILED", "OFFBOARDING", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseChromeOSOnboardingStatus(v string) (interface{}, error) {
    result := UNKNOWN_CHROMEOSONBOARDINGSTATUS
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_CHROMEOSONBOARDINGSTATUS
        case "INPROGRESS":
            result = INPROGRESS_CHROMEOSONBOARDINGSTATUS
        case "ONBOARDED":
            result = ONBOARDED_CHROMEOSONBOARDINGSTATUS
        case "FAILED":
            result = FAILED_CHROMEOSONBOARDINGSTATUS
        case "OFFBOARDING":
            result = OFFBOARDING_CHROMEOSONBOARDINGSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_CHROMEOSONBOARDINGSTATUS
        default:
            return 0, errors.New("Unknown ChromeOSOnboardingStatus value: " + v)
    }
    return &result, nil
}
func SerializeChromeOSOnboardingStatus(values []ChromeOSOnboardingStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
