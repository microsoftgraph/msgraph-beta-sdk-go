package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "STANDARD":
            return STANDARD_ASSIGNMENTMETHOD, nil
        case "PRIVILEGED":
            return PRIVILEGED_ASSIGNMENTMETHOD, nil
        case "AUTO":
            return AUTO_ASSIGNMENTMETHOD, nil
    }
    return 0, errors.New("Unknown AssignmentMethod value: " + v)
}
func SerializeAssignmentMethod(values []AssignmentMethod) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
