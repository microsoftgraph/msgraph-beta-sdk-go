package graph
import (
    "strings"
    "errors"
)
// 
type AssignmentMethod int

const (
    STANDARD_ASSIGNMENTMETHOD AssignmentMethod = iota
    PRIVILEGED_ASSIGNMENTMETHOD
    AUTO_ASSIGNMENTMETHOD
)

func (i AssignmentMethod) String() string {
    return []string{"STANDARD", "PRIVILEGED", "AUTO"}[i]
}
func ParseAssignmentMethod(v string) (interface{}, error) {
    result := STANDARD_ASSIGNMENTMETHOD
    switch strings.ToUpper(v) {
        case "STANDARD":
            result = STANDARD_ASSIGNMENTMETHOD
        case "PRIVILEGED":
            result = PRIVILEGED_ASSIGNMENTMETHOD
        case "AUTO":
            result = AUTO_ASSIGNMENTMETHOD
        default:
            return 0, errors.New("Unknown AssignmentMethod value: " + v)
    }
    return &result, nil
}
func SerializeAssignmentMethod(values []AssignmentMethod) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
