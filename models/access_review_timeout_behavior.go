package models
import (
    "errors"
)
type AccessReviewTimeoutBehavior int

const (
    KEEPACCESS_ACCESSREVIEWTIMEOUTBEHAVIOR AccessReviewTimeoutBehavior = iota
    REMOVEACCESS_ACCESSREVIEWTIMEOUTBEHAVIOR
    ACCEPTACCESSRECOMMENDATION_ACCESSREVIEWTIMEOUTBEHAVIOR
    UNKNOWNFUTUREVALUE_ACCESSREVIEWTIMEOUTBEHAVIOR
)

func (i AccessReviewTimeoutBehavior) String() string {
    return []string{"keepAccess", "removeAccess", "acceptAccessRecommendation", "unknownFutureValue"}[i]
}
func ParseAccessReviewTimeoutBehavior(v string) (any, error) {
    result := KEEPACCESS_ACCESSREVIEWTIMEOUTBEHAVIOR
    switch v {
        case "keepAccess":
            result = KEEPACCESS_ACCESSREVIEWTIMEOUTBEHAVIOR
        case "removeAccess":
            result = REMOVEACCESS_ACCESSREVIEWTIMEOUTBEHAVIOR
        case "acceptAccessRecommendation":
            result = ACCEPTACCESSRECOMMENDATION_ACCESSREVIEWTIMEOUTBEHAVIOR
        case "unknownFutureValue":
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
func (i AccessReviewTimeoutBehavior) isMultiValue() bool {
    return false
}
