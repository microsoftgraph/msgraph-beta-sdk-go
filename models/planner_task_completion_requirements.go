package models
import (
    "errors"
    "math"
    "strings"
)
type PlannerTaskCompletionRequirements int

const (
    NONE_PLANNERTASKCOMPLETIONREQUIREMENTS = 1
    CHECKLISTCOMPLETION_PLANNERTASKCOMPLETIONREQUIREMENTS = 2
    UNKNOWNFUTUREVALUE_PLANNERTASKCOMPLETIONREQUIREMENTS = 4
    FORMCOMPLETION_PLANNERTASKCOMPLETIONREQUIREMENTS = 8
    APPROVALCOMPLETION_PLANNERTASKCOMPLETIONREQUIREMENTS = 16
)

func (i PlannerTaskCompletionRequirements) String() string {
    var values []string
    options := []string{"none", "checklistCompletion", "unknownFutureValue", "formCompletion", "approvalCompletion"}
    for p := 0; p < 5; p++ {
        mantis := PlannerTaskCompletionRequirements(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
        }
    }
    return strings.Join(values, ",")
}
func ParsePlannerTaskCompletionRequirements(v string) (any, error) {
    var result PlannerTaskCompletionRequirements
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_PLANNERTASKCOMPLETIONREQUIREMENTS
            case "checklistCompletion":
                result |= CHECKLISTCOMPLETION_PLANNERTASKCOMPLETIONREQUIREMENTS
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_PLANNERTASKCOMPLETIONREQUIREMENTS
            case "formCompletion":
                result |= FORMCOMPLETION_PLANNERTASKCOMPLETIONREQUIREMENTS
            case "approvalCompletion":
                result |= APPROVALCOMPLETION_PLANNERTASKCOMPLETIONREQUIREMENTS
            default:
                return 0, errors.New("Unknown PlannerTaskCompletionRequirements value: " + v)
        }
    }
    return &result, nil
}
func SerializePlannerTaskCompletionRequirements(values []PlannerTaskCompletionRequirements) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PlannerTaskCompletionRequirements) isMultiValue() bool {
    return true
}
