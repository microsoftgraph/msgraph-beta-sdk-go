package graph
import (
    "strings"
    "errors"
)
// 
type AccessReviewTimeoutBehavior int

const (
    KEEPACCESS_ACCESSREVIEWTIMEOUTBEHAVIOR AccessReviewTimeoutBehavior = iota
    REMOVEACCESS_ACCESSREVIEWTIMEOUTBEHAVIOR
    ACCEPTACCESSRECOMMENDATION_ACCESSREVIEWTIMEOUTBEHAVIOR
    UNKNOWNFUTUREVALUE_ACCESSREVIEWTIMEOUTBEHAVIOR
)

func (i AccessReviewTimeoutBehavior) String() string {
    return []string{"KEEPACCESS", "REMOVEACCESS", "ACCEPTACCESSRECOMMENDATION", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAccessReviewTimeoutBehavior(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "KEEPACCESS":
            return KEEPACCESS_ACCESSREVIEWTIMEOUTBEHAVIOR, nil
        case "REMOVEACCESS":
            return REMOVEACCESS_ACCESSREVIEWTIMEOUTBEHAVIOR, nil
        case "ACCEPTACCESSRECOMMENDATION":
            return ACCEPTACCESSRECOMMENDATION_ACCESSREVIEWTIMEOUTBEHAVIOR, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_ACCESSREVIEWTIMEOUTBEHAVIOR, nil
    }
    return 0, errors.New("Unknown AccessReviewTimeoutBehavior value: " + v)
}
func SerializeAccessReviewTimeoutBehavior(values []AccessReviewTimeoutBehavior) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
