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
    switch strings.ToUpper(v) {
        case "PUBLISHED":
            return PUBLISHED_ANSWERSTATE, nil
        case "DRAFT":
            return DRAFT_ANSWERSTATE, nil
        case "EXCLUDED":
            return EXCLUDED_ANSWERSTATE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_ANSWERSTATE, nil
    }
    return 0, errors.New("Unknown AnswerState value: " + v)
}
func SerializeAnswerState(values []AnswerState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
