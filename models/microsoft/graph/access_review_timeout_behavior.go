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
    result := KEEPACCESS_ACCESSREVIEWTIMEOUTBEHAVIOR
    switch strings.ToUpper(v) {
        case "KEEPACCESS":
            result = KEEPACCESS_ACCESSREVIEWTIMEOUTBEHAVIOR
        case "REMOVEACCESS":
            result = REMOVEACCESS_ACCESSREVIEWTIMEOUTBEHAVIOR
        case "ACCEPTACCESSRECOMMENDATION":
            result = ACCEPTACCESSRECOMMENDATION_ACCESSREVIEWTIMEOUTBEHAVIOR
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ACCESSREVIEWTIMEOUTBEHAVIOR
        default:
            return 0, errors.New("Unknown AccessReviewTimeoutBehavior value: " + v)
    }
    return &result, nil
}
func SerializeAccessReviewTimeoutBehavior(values []AccessReviewTimeoutBehavior) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
