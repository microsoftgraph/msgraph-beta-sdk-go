package models
import (
    "errors"
    "strings"
)
// 
type PlannerTaskCompletionRequirements int

const (
    NONE_PLANNERTASKCOMPLETIONREQUIREMENTS PlannerTaskCompletionRequirements = iota
    CHECKLISTCOMPLETION_PLANNERTASKCOMPLETIONREQUIREMENTS
    UNKNOWNFUTUREVALUE_PLANNERTASKCOMPLETIONREQUIREMENTS
)

func (i PlannerTaskCompletionRequirements) String() string {
    var values []string
    for p := PlannerTaskCompletionRequirements(1); p <= UNKNOWNFUTUREVALUE_PLANNERTASKCOMPLETIONREQUIREMENTS; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"none", "checklistCompletion", "unknownFutureValue"}[p])
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
