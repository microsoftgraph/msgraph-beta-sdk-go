package models
import (
    "errors"
)
// An admin enforced Android SafetyNet evaluation type requirement on a managed app.
type AndroidManagedAppSafetyNetEvaluationType int

const (
    // Require basic evaluation
    BASIC_ANDROIDMANAGEDAPPSAFETYNETEVALUATIONTYPE AndroidManagedAppSafetyNetEvaluationType = iota
    // Require hardware backed evaluation
    HARDWAREBACKED_ANDROIDMANAGEDAPPSAFETYNETEVALUATIONTYPE
)

func (i AndroidManagedAppSafetyNetEvaluationType) String() string {
    return []string{"basic", "hardwareBacked"}[i]
}
func ParseAndroidManagedAppSafetyNetEvaluationType(v string) (any, error) {
    result := BASIC_ANDROIDMANAGEDAPPSAFETYNETEVALUATIONTYPE
    switch v {
        case "basic":
            result = BASIC_ANDROIDMANAGEDAPPSAFETYNETEVALUATIONTYPE
        case "hardwareBacked":
            result = HARDWAREBACKED_ANDROIDMANAGEDAPPSAFETYNETEVALUATIONTYPE
        default:
            return 0, errors.New("Unknown AndroidManagedAppSafetyNetEvaluationType value: " + v)
    }
    return &result, nil
}
func SerializeAndroidManagedAppSafetyNetEvaluationType(values []AndroidManagedAppSafetyNetEvaluationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AndroidManagedAppSafetyNetEvaluationType) isMultiValue() bool {
    return false
}
