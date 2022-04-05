package models
import (
    "strings"
    "errors"
)
// Provides operations to call the extractSensitivityLabels method.
type SensitivityLabelAssignmentMethod int

const (
    STANDARD_SENSITIVITYLABELASSIGNMENTMETHOD SensitivityLabelAssignmentMethod = iota
    PRIVILEGED_SENSITIVITYLABELASSIGNMENTMETHOD
    AUTO_SENSITIVITYLABELASSIGNMENTMETHOD
    UNKNOWNFUTUREVALUE_SENSITIVITYLABELASSIGNMENTMETHOD
)

func (i SensitivityLabelAssignmentMethod) String() string {
    return []string{"STANDARD", "PRIVILEGED", "AUTO", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseSensitivityLabelAssignmentMethod(v string) (interface{}, error) {
    result := STANDARD_SENSITIVITYLABELASSIGNMENTMETHOD
    switch strings.ToUpper(v) {
        case "STANDARD":
            result = STANDARD_SENSITIVITYLABELASSIGNMENTMETHOD
        case "PRIVILEGED":
            result = PRIVILEGED_SENSITIVITYLABELASSIGNMENTMETHOD
        case "AUTO":
            result = AUTO_SENSITIVITYLABELASSIGNMENTMETHOD
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_SENSITIVITYLABELASSIGNMENTMETHOD
        default:
            return 0, errors.New("Unknown SensitivityLabelAssignmentMethod value: " + v)
    }
    return &result, nil
}
func SerializeSensitivityLabelAssignmentMethod(values []SensitivityLabelAssignmentMethod) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
