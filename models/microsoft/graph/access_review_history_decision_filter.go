package graph
import (
    "strings"
    "errors"
)
// 
type AccessReviewHistoryDecisionFilter int

const (
    APPROVE_ACCESSREVIEWHISTORYDECISIONFILTER AccessReviewHistoryDecisionFilter = iota
    DENY_ACCESSREVIEWHISTORYDECISIONFILTER
    NOTREVIEWED_ACCESSREVIEWHISTORYDECISIONFILTER
    DONTKNOW_ACCESSREVIEWHISTORYDECISIONFILTER
    NOTNOTIFIED_ACCESSREVIEWHISTORYDECISIONFILTER
    UNKNOWNFUTUREVALUE_ACCESSREVIEWHISTORYDECISIONFILTER
)

func (i AccessReviewHistoryDecisionFilter) String() string {
    return []string{"APPROVE", "DENY", "NOTREVIEWED", "DONTKNOW", "NOTNOTIFIED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAccessReviewHistoryDecisionFilter(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "APPROVE":
            return APPROVE_ACCESSREVIEWHISTORYDECISIONFILTER, nil
        case "DENY":
            return DENY_ACCESSREVIEWHISTORYDECISIONFILTER, nil
        case "NOTREVIEWED":
            return NOTREVIEWED_ACCESSREVIEWHISTORYDECISIONFILTER, nil
        case "DONTKNOW":
            return DONTKNOW_ACCESSREVIEWHISTORYDECISIONFILTER, nil
        case "NOTNOTIFIED":
            return NOTNOTIFIED_ACCESSREVIEWHISTORYDECISIONFILTER, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_ACCESSREVIEWHISTORYDECISIONFILTER, nil
    }
    return 0, errors.New("Unknown AccessReviewHistoryDecisionFilter value: " + v)
}
func SerializeAccessReviewHistoryDecisionFilter(values []AccessReviewHistoryDecisionFilter) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
