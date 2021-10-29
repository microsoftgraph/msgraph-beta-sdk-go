package graph
import (
    "strings"
    "errors"
)
// 
type AccessReviewHistoryStatus int

const (
    DONE_ACCESSREVIEWHISTORYSTATUS AccessReviewHistoryStatus = iota
    INPROGRESS_ACCESSREVIEWHISTORYSTATUS
    ERROR_ACCESSREVIEWHISTORYSTATUS
    REQUESTED_ACCESSREVIEWHISTORYSTATUS
    UNKNOWNFUTUREVALUE_ACCESSREVIEWHISTORYSTATUS
)

func (i AccessReviewHistoryStatus) String() string {
    return []string{"DONE", "INPROGRESS", "ERROR", "REQUESTED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAccessReviewHistoryStatus(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "DONE":
            return DONE_ACCESSREVIEWHISTORYSTATUS, nil
        case "INPROGRESS":
            return INPROGRESS_ACCESSREVIEWHISTORYSTATUS, nil
        case "ERROR":
            return ERROR_ACCESSREVIEWHISTORYSTATUS, nil
        case "REQUESTED":
            return REQUESTED_ACCESSREVIEWHISTORYSTATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_ACCESSREVIEWHISTORYSTATUS, nil
    }
    return 0, errors.New("Unknown AccessReviewHistoryStatus value: " + v)
}
func SerializeAccessReviewHistoryStatus(values []AccessReviewHistoryStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
