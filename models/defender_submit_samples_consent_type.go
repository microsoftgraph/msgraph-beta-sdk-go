package models
import (
    "errors"
)
// Possible values for DefenderSubmitSamplesConsentType
type DefenderSubmitSamplesConsentType int

const (
    // Send safe samples automatically
    SENDSAFESAMPLESAUTOMATICALLY_DEFENDERSUBMITSAMPLESCONSENTTYPE DefenderSubmitSamplesConsentType = iota
    // Always prompt
    ALWAYSPROMPT_DEFENDERSUBMITSAMPLESCONSENTTYPE
    // Never send
    NEVERSEND_DEFENDERSUBMITSAMPLESCONSENTTYPE
    // Send all samples automatically
    SENDALLSAMPLESAUTOMATICALLY_DEFENDERSUBMITSAMPLESCONSENTTYPE
)

func (i DefenderSubmitSamplesConsentType) String() string {
    return []string{"sendSafeSamplesAutomatically", "alwaysPrompt", "neverSend", "sendAllSamplesAutomatically"}[i]
}
func ParseDefenderSubmitSamplesConsentType(v string) (any, error) {
    result := SENDSAFESAMPLESAUTOMATICALLY_DEFENDERSUBMITSAMPLESCONSENTTYPE
    switch v {
        case "sendSafeSamplesAutomatically":
            result = SENDSAFESAMPLESAUTOMATICALLY_DEFENDERSUBMITSAMPLESCONSENTTYPE
        case "alwaysPrompt":
            result = ALWAYSPROMPT_DEFENDERSUBMITSAMPLESCONSENTTYPE
        case "neverSend":
            result = NEVERSEND_DEFENDERSUBMITSAMPLESCONSENTTYPE
        case "sendAllSamplesAutomatically":
            result = SENDALLSAMPLESAUTOMATICALLY_DEFENDERSUBMITSAMPLESCONSENTTYPE
        default:
            return 0, errors.New("Unknown DefenderSubmitSamplesConsentType value: " + v)
    }
    return &result, nil
}
func SerializeDefenderSubmitSamplesConsentType(values []DefenderSubmitSamplesConsentType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DefenderSubmitSamplesConsentType) isMultiValue() bool {
    return false
}
