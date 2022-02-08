package graph
import (
    "strings"
    "errors"
)
// 
type AndroidManagedAppSafetyNetEvaluationType int

const (
    BASIC_ANDROIDMANAGEDAPPSAFETYNETEVALUATIONTYPE AndroidManagedAppSafetyNetEvaluationType = iota
    HARDWAREBACKED_ANDROIDMANAGEDAPPSAFETYNETEVALUATIONTYPE
)

func (i AndroidManagedAppSafetyNetEvaluationType) String() string {
    return []string{"BASIC", "HARDWAREBACKED"}[i]
}
func ParseAndroidManagedAppSafetyNetEvaluationType(v string) (interface{}, error) {
    result := BASIC_ANDROIDMANAGEDAPPSAFETYNETEVALUATIONTYPE
    switch strings.ToUpper(v) {
        case "BASIC":
            result = BASIC_ANDROIDMANAGEDAPPSAFETYNETEVALUATIONTYPE
        case "HARDWAREBACKED":
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
