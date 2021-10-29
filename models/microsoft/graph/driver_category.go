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
    switch strings.ToUpper(v) {
        case "RECOMMENDED":
            return RECOMMENDED_DRIVERCATEGORY, nil
        case "PREVIOUSLYAPPROVED":
            return PREVIOUSLYAPPROVED_DRIVERCATEGORY, nil
        case "OTHER":
            return OTHER_DRIVERCATEGORY, nil
    }
    return 0, errors.New("Unknown DriverCategory value: " + v)
}
func SerializeDriverCategory(values []DriverCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
