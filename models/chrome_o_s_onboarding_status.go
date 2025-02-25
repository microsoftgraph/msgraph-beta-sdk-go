package models
// The onboarding status of the tenant.
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
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_CHROMEOSONBOARDINGSTATUS
)

func (i ChromeOSOnboardingStatus) String() string {
    return []string{"unknown", "inprogress", "onboarded", "failed", "offboarding", "unknownFutureValue"}[i]
}
func ParseChromeOSOnboardingStatus(v string) (any, error) {
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
            return nil, nil
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
func (i ChromeOSOnboardingStatus) isMultiValue() bool {
    return false
}
