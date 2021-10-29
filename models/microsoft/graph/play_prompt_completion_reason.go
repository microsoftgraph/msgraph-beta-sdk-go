package graph
import (
    "strings"
    "errors"
)
// 
type PlayPromptCompletionReason int

const (
    UNKNOWN_PLAYPROMPTCOMPLETIONREASON PlayPromptCompletionReason = iota
    COMPLETEDSUCCESSFULLY_PLAYPROMPTCOMPLETIONREASON
    MEDIAOPERATIONCANCELED_PLAYPROMPTCOMPLETIONREASON
    UNKNOWNFUTUREVALUE_PLAYPROMPTCOMPLETIONREASON
)

func (i PlayPromptCompletionReason) String() string {
    return []string{"UNKNOWN", "COMPLETEDSUCCESSFULLY", "MEDIAOPERATIONCANCELED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParsePlayPromptCompletionReason(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_PLAYPROMPTCOMPLETIONREASON, nil
        case "COMPLETEDSUCCESSFULLY":
            return COMPLETEDSUCCESSFULLY_PLAYPROMPTCOMPLETIONREASON, nil
        case "MEDIAOPERATIONCANCELED":
            return MEDIAOPERATIONCANCELED_PLAYPROMPTCOMPLETIONREASON, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_PLAYPROMPTCOMPLETIONREASON, nil
    }
    return 0, errors.New("Unknown PlayPromptCompletionReason value: " + v)
}
func SerializePlayPromptCompletionReason(values []PlayPromptCompletionReason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
