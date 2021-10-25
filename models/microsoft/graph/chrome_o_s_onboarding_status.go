package graph
import (
    "strings"
    "errors"
)
type ChromeOSOnboardingStatus int

const (
    UNKNOWN_CHROMEOSONBOARDINGSTATUS ChromeOSOnboardingStatus = iota
    INPROGRESS_CHROMEOSONBOARDINGSTATUS
    ONBOARDED_CHROMEOSONBOARDINGSTATUS
    FAILED_CHROMEOSONBOARDINGSTATUS
)

func (i ChromeOSOnboardingStatus) String() string {
    return []string{"UNKNOWN", "INPROGRESS", "ONBOARDED", "FAILED"}[i]
}
func ParseChromeOSOnboardingStatus(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_CHROMEOSONBOARDINGSTATUS, nil
        case "INPROGRESS":
            return INPROGRESS_CHROMEOSONBOARDINGSTATUS, nil
        case "ONBOARDED":
            return ONBOARDED_CHROMEOSONBOARDINGSTATUS, nil
        case "FAILED":
            return FAILED_CHROMEOSONBOARDINGSTATUS, nil
    }
    return 0, errors.New("Unknown ChromeOSOnboardingStatus value: " + v)
}
func SerializeChromeOSOnboardingStatus(values []ChromeOSOnboardingStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
