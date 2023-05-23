package models
import (
    "errors"
)
// 
type PlannerTaskCompletionRequirements int

const (
    NONE_PLANNERTASKCOMPLETIONREQUIREMENTS PlannerTaskCompletionRequirements = iota
    CHECKLISTCOMPLETION_PLANNERTASKCOMPLETIONREQUIREMENTS
    UNKNOWNFUTUREVALUE_PLANNERTASKCOMPLETIONREQUIREMENTS
)

func (i PlannerTaskCompletionRequirements) String() string {
    return []string{"none", "checklistCompletion", "unknownFutureValue"}[i]
}
func ParsePlannerTaskCompletionRequirements(v string) (any, error) {
    result := NONE_PLANNERTASKCOMPLETIONREQUIREMENTS
    switch v {
        case "none":
            result = NONE_PLANNERTASKCOMPLETIONREQUIREMENTS
        case "checklistCompletion":
            result = CHECKLISTCOMPLETION_PLANNERTASKCOMPLETIONREQUIREMENTS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PLANNERTASKCOMPLETIONREQUIREMENTS
        default:
            return 0, errors.New("Unknown PlannerTaskCompletionRequirements value: " + v)
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
