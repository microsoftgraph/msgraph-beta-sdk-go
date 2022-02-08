package graph
import (
    "strings"
    "errors"
)
// 
type DriverCategory int

const (
    RECOMMENDED_DRIVERCATEGORY DriverCategory = iota
    PREVIOUSLYAPPROVED_DRIVERCATEGORY
    OTHER_DRIVERCATEGORY
)

func (i DriverCategory) String() string {
    return []string{"RECOMMENDED", "PREVIOUSLYAPPROVED", "OTHER"}[i]
}
func ParseDriverCategory(v string) (interface{}, error) {
    result := RECOMMENDED_DRIVERCATEGORY
    switch strings.ToUpper(v) {
        case "RECOMMENDED":
            result = RECOMMENDED_DRIVERCATEGORY
        case "PREVIOUSLYAPPROVED":
            result = PREVIOUSLYAPPROVED_DRIVERCATEGORY
        case "OTHER":
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
