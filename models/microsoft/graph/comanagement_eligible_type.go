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
    switch strings.ToUpper(v) {
        case "COMANAGED":
            return COMANAGED_COMANAGEMENTELIGIBLETYPE, nil
        case "ELIGIBLE":
            return ELIGIBLE_COMANAGEMENTELIGIBLETYPE, nil
        case "ELIGIBLEBUTNOTAZUREADJOINED":
            return ELIGIBLEBUTNOTAZUREADJOINED_COMANAGEMENTELIGIBLETYPE, nil
        case "NEEDSOSUPDATE":
            return NEEDSOSUPDATE_COMANAGEMENTELIGIBLETYPE, nil
        case "INELIGIBLE":
            return INELIGIBLE_COMANAGEMENTELIGIBLETYPE, nil
    }
    return 0, errors.New("Unknown ComanagementEligibleType value: " + v)
}
func SerializeComanagementEligibleType(values []ComanagementEligibleType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
