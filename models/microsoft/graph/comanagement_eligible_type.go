package graph
import (
    "strings"
    "errors"
)
// 
type ComanagementEligibleType int

const (
    COMANAGED_COMANAGEMENTELIGIBLETYPE ComanagementEligibleType = iota
    ELIGIBLE_COMANAGEMENTELIGIBLETYPE
    ELIGIBLEBUTNOTAZUREADJOINED_COMANAGEMENTELIGIBLETYPE
    NEEDSOSUPDATE_COMANAGEMENTELIGIBLETYPE
    INELIGIBLE_COMANAGEMENTELIGIBLETYPE
)

func (i ComanagementEligibleType) String() string {
    return []string{"COMANAGED", "ELIGIBLE", "ELIGIBLEBUTNOTAZUREADJOINED", "NEEDSOSUPDATE", "INELIGIBLE"}[i]
}
func ParseComanagementEligibleType(v string) (interface{}, error) {
    result := COMANAGED_COMANAGEMENTELIGIBLETYPE
    switch strings.ToUpper(v) {
        case "COMANAGED":
            result = COMANAGED_COMANAGEMENTELIGIBLETYPE
        case "ELIGIBLE":
            result = ELIGIBLE_COMANAGEMENTELIGIBLETYPE
        case "ELIGIBLEBUTNOTAZUREADJOINED":
            result = ELIGIBLEBUTNOTAZUREADJOINED_COMANAGEMENTELIGIBLETYPE
        case "NEEDSOSUPDATE":
            result = NEEDSOSUPDATE_COMANAGEMENTELIGIBLETYPE
        case "INELIGIBLE":
            result = INELIGIBLE_COMANAGEMENTELIGIBLETYPE
        default:
            return 0, errors.New("Unknown ComanagementEligibleType value: " + v)
    }
    return &result, nil
}
func SerializeComanagementEligibleType(values []ComanagementEligibleType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
