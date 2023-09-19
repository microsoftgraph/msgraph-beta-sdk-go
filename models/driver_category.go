package models
import (
    "errors"
)
// An enum type to represent which category a driver belongs to.
type DriverCategory int

const (
    // This indicates a driver is recommended by Microsoft.
    RECOMMENDED_DRIVERCATEGORY DriverCategory = iota
    // This indicates a driver was recommended by Microsoft and IT admin has taken some approval action on it.
    PREVIOUSLYAPPROVED_DRIVERCATEGORY
    // This indicates a driver is never recommended by Microsoft.
    OTHER_DRIVERCATEGORY
)

func (i DriverCategory) String() string {
    return []string{"recommended", "previouslyApproved", "other"}[i]
}
func ParseDriverCategory(v string) (any, error) {
    result := RECOMMENDED_DRIVERCATEGORY
    switch v {
        case "recommended":
            result = RECOMMENDED_DRIVERCATEGORY
        case "previouslyApproved":
            result = PREVIOUSLYAPPROVED_DRIVERCATEGORY
        case "other":
            result = OTHER_DRIVERCATEGORY
        default:
            return 0, errors.New("Unknown DriverCategory value: " + v)
    }
    return &result, nil
}
func SerializeDriverCategory(values []DriverCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DriverCategory) isMultiValue() bool {
    return false
}
