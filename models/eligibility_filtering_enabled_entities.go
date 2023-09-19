package models
import (
    "errors"
    "strings"
)
// 
type EligibilityFilteringEnabledEntities int

const (
    NONE_ELIGIBILITYFILTERINGENABLEDENTITIES EligibilityFilteringEnabledEntities = iota
    SWAPREQUEST_ELIGIBILITYFILTERINGENABLEDENTITIES
    OFFERSHIFTREQUEST_ELIGIBILITYFILTERINGENABLEDENTITIES
    UNKNOWNFUTUREVALUE_ELIGIBILITYFILTERINGENABLEDENTITIES
    TIMEOFFREASON_ELIGIBILITYFILTERINGENABLEDENTITIES
)

func (i EligibilityFilteringEnabledEntities) String() string {
    var values []string
    for p := EligibilityFilteringEnabledEntities(1); p <= TIMEOFFREASON_ELIGIBILITYFILTERINGENABLEDENTITIES; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"none", "swapRequest", "offerShiftRequest", "unknownFutureValue", "timeOffReason"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseEligibilityFilteringEnabledEntities(v string) (any, error) {
    var result EligibilityFilteringEnabledEntities
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_ELIGIBILITYFILTERINGENABLEDENTITIES
            case "swapRequest":
                result |= SWAPREQUEST_ELIGIBILITYFILTERINGENABLEDENTITIES
            case "offerShiftRequest":
                result |= OFFERSHIFTREQUEST_ELIGIBILITYFILTERINGENABLEDENTITIES
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_ELIGIBILITYFILTERINGENABLEDENTITIES
            case "timeOffReason":
                result |= TIMEOFFREASON_ELIGIBILITYFILTERINGENABLEDENTITIES
            default:
                return 0, errors.New("Unknown EligibilityFilteringEnabledEntities value: " + v)
        }
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
func (i EligibilityFilteringEnabledEntities) isMultiValue() bool {
    return true
}
