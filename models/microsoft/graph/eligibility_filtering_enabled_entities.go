package graph
import (
    "strings"
    "errors"
)
// 
type EligibilityFilteringEnabledEntities int

const (
    NONE_ELIGIBILITYFILTERINGENABLEDENTITIES EligibilityFilteringEnabledEntities = iota
    SWAPREQUEST_ELIGIBILITYFILTERINGENABLEDENTITIES
    OFFERSHIFTREQUEST_ELIGIBILITYFILTERINGENABLEDENTITIES
    UNKNOWNFUTUREVALUE_ELIGIBILITYFILTERINGENABLEDENTITIES
)

func (i EligibilityFilteringEnabledEntities) String() string {
    return []string{"NONE", "SWAPREQUEST", "OFFERSHIFTREQUEST", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseEligibilityFilteringEnabledEntities(v string) (interface{}, error) {
    result := NONE_ELIGIBILITYFILTERINGENABLEDENTITIES
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_ELIGIBILITYFILTERINGENABLEDENTITIES
        case "SWAPREQUEST":
            result = SWAPREQUEST_ELIGIBILITYFILTERINGENABLEDENTITIES
        case "OFFERSHIFTREQUEST":
            result = OFFERSHIFTREQUEST_ELIGIBILITYFILTERINGENABLEDENTITIES
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ELIGIBILITYFILTERINGENABLEDENTITIES
        default:
            return 0, errors.New("Unknown EligibilityFilteringEnabledEntities value: " + v)
    }
    return &result, nil
}
func SerializeEligibilityFilteringEnabledEntities(values []EligibilityFilteringEnabledEntities) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
