package search
import (
    "strings"
    "errors"
)
// 
type AnswerState int

const (
    PUBLISHED_ANSWERSTATE AnswerState = iota
    DRAFT_ANSWERSTATE
    EXCLUDED_ANSWERSTATE
    UNKNOWNFUTUREVALUE_ANSWERSTATE
)

func (i AnswerState) String() string {
    return []string{"PUBLISHED", "DRAFT", "EXCLUDED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAnswerState(v string) (interface{}, error) {
    result := PUBLISHED_ANSWERSTATE
    switch strings.ToUpper(v) {
        case "PUBLISHED":
            result = PUBLISHED_ANSWERSTATE
        case "DRAFT":
            result = DRAFT_ANSWERSTATE
        case "EXCLUDED":
            result = EXCLUDED_ANSWERSTATE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ANSWERSTATE
        default:
            return 0, errors.New("Unknown AnswerState value: " + v)
    }
    return &result, nil
}
func SerializeAnswerState(values []AnswerState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
