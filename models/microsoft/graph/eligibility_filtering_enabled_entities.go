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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_ELIGIBILITYFILTERINGENABLEDENTITIES, nil
        case "SWAPREQUEST":
            return SWAPREQUEST_ELIGIBILITYFILTERINGENABLEDENTITIES, nil
        case "OFFERSHIFTREQUEST":
            return OFFERSHIFTREQUEST_ELIGIBILITYFILTERINGENABLEDENTITIES, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_ELIGIBILITYFILTERINGENABLEDENTITIES, nil
    }
    return 0, errors.New("Unknown EligibilityFilteringEnabledEntities value: " + v)
}
func SerializeEligibilityFilteringEnabledEntities(values []EligibilityFilteringEnabledEntities) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
