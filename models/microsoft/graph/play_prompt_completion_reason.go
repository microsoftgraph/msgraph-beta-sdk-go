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
    result := UNKNOWN_PLAYPROMPTCOMPLETIONREASON
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_PLAYPROMPTCOMPLETIONREASON
        case "COMPLETEDSUCCESSFULLY":
            result = COMPLETEDSUCCESSFULLY_PLAYPROMPTCOMPLETIONREASON
        case "MEDIAOPERATIONCANCELED":
            result = MEDIAOPERATIONCANCELED_PLAYPROMPTCOMPLETIONREASON
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_PLAYPROMPTCOMPLETIONREASON
        default:
            return 0, errors.New("Unknown PlayPromptCompletionReason value: " + v)
    }
    return &result, nil
}
func SerializePlayPromptCompletionReason(values []PlayPromptCompletionReason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
