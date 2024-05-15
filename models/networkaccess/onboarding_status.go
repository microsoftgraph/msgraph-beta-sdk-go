package networkaccess
type OnboardingStatus int

const (
    OFFBOARDED_ONBOARDINGSTATUS OnboardingStatus = iota
    OFFBOARDINGINPROGRESS_ONBOARDINGSTATUS
    ONBOARDINGINPROGRESS_ONBOARDINGSTATUS
    ONBOARDED_ONBOARDINGSTATUS
    ONBOARDINGERROROCCURRED_ONBOARDINGSTATUS
    OFFBOARDINGERROROCCURRED_ONBOARDINGSTATUS
    UNKNOWNFUTUREVALUE_ONBOARDINGSTATUS
)

func (i OnboardingStatus) String() string {
    return []string{"offboarded", "offboardingInProgress", "onboardingInProgress", "onboarded", "onboardingErrorOccurred", "offboardingErrorOccurred", "unknownFutureValue"}[i]
}
func ParseOnboardingStatus(v string) (any, error) {
    result := OFFBOARDED_ONBOARDINGSTATUS
    switch v {
        case "offboarded":
            result = OFFBOARDED_ONBOARDINGSTATUS
        case "offboardingInProgress":
            result = OFFBOARDINGINPROGRESS_ONBOARDINGSTATUS
        case "onboardingInProgress":
            result = ONBOARDINGINPROGRESS_ONBOARDINGSTATUS
        case "onboarded":
            result = ONBOARDED_ONBOARDINGSTATUS
        case "onboardingErrorOccurred":
            result = ONBOARDINGERROROCCURRED_ONBOARDINGSTATUS
        case "offboardingErrorOccurred":
            result = OFFBOARDINGERROROCCURRED_ONBOARDINGSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ONBOARDINGSTATUS
        default:
            return nil, nil
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
func (i OnboardingStatus) isMultiValue() bool {
    return false
}
