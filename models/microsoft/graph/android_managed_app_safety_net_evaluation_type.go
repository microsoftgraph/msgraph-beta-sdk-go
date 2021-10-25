package graph
import (
    "strings"
    "errors"
)
type AndroidManagedAppSafetyNetEvaluationType int

const (
    BASIC_ANDROIDMANAGEDAPPSAFETYNETEVALUATIONTYPE AndroidManagedAppSafetyNetEvaluationType = iota
    HARDWAREBACKED_ANDROIDMANAGEDAPPSAFETYNETEVALUATIONTYPE
)

func (i AndroidManagedAppSafetyNetEvaluationType) String() string {
    return []string{"BASIC", "HARDWAREBACKED"}[i]
}
func ParseAndroidManagedAppSafetyNetEvaluationType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "BASIC":
            return BASIC_ANDROIDMANAGEDAPPSAFETYNETEVALUATIONTYPE, nil
        case "HARDWAREBACKED":
            return HARDWAREBACKED_ANDROIDMANAGEDAPPSAFETYNETEVALUATIONTYPE, nil
    }
    return 0, errors.New("Unknown AndroidManagedAppSafetyNetEvaluationType value: " + v)
}
func SerializeAndroidManagedAppSafetyNetEvaluationType(values []AndroidManagedAppSafetyNetEvaluationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
