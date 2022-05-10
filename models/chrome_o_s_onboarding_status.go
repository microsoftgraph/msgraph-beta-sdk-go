package models
import (
    "errors"
)
// Provides operations to call the connect method.
type ChromeOSOnboardingStatus int

const (
    // Unknown
    UNKNOWN_CHROMEOSONBOARDINGSTATUS ChromeOSOnboardingStatus = iota
    // In progress
    INPROGRESS_CHROMEOSONBOARDINGSTATUS
    // Onboarded
    ONBOARDED_CHROMEOSONBOARDINGSTATUS
    // Failed
    FAILED_CHROMEOSONBOARDINGSTATUS
    // Offboarding
    OFFBOARDING_CHROMEOSONBOARDINGSTATUS
    // UnknownFutureValue
    UNKNOWNFUTUREVALUE_CHROMEOSONBOARDINGSTATUS
)

func (i ChromeOSOnboardingStatus) String() string {
    return []string{"unknown", "inprogress", "onboarded", "failed", "offboarding", "unknownFutureValue"}[i]
}
func ParseChromeOSOnboardingStatus(v string) (interface{}, error) {
    result := UNKNOWN_CHROMEOSONBOARDINGSTATUS
    switch v {
        case "unknown":
            result = UNKNOWN_CHROMEOSONBOARDINGSTATUS
        case "inprogress":
            result = INPROGRESS_CHROMEOSONBOARDINGSTATUS
        case "onboarded":
            result = ONBOARDED_CHROMEOSONBOARDINGSTATUS
        case "failed":
            result = FAILED_CHROMEOSONBOARDINGSTATUS
        case "offboarding":
            result = OFFBOARDING_CHROMEOSONBOARDINGSTATUS
        case "unknownFutureValue":
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
