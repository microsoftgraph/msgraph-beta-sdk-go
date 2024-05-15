package models
// Supported filter management types whether its devices or apps.
type AssignmentFilterManagementType int

const (
    // Indicates when filter is supported based on device properties. This is the default value when management type resolution fails.
    DEVICES_ASSIGNMENTFILTERMANAGEMENTTYPE AssignmentFilterManagementType = iota
    // Indicates when filter is supported based on app properties.
    APPS_ASSIGNMENTFILTERMANAGEMENTTYPE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_ASSIGNMENTFILTERMANAGEMENTTYPE
)

func (i AssignmentFilterManagementType) String() string {
    return []string{"devices", "apps", "unknownFutureValue"}[i]
}
func ParseAssignmentFilterManagementType(v string) (any, error) {
    result := DEVICES_ASSIGNMENTFILTERMANAGEMENTTYPE
    switch v {
        case "devices":
            result = DEVICES_ASSIGNMENTFILTERMANAGEMENTTYPE
        case "apps":
            result = APPS_ASSIGNMENTFILTERMANAGEMENTTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ASSIGNMENTFILTERMANAGEMENTTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAssignmentFilterManagementType(values []AssignmentFilterManagementType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AssignmentFilterManagementType) isMultiValue() bool {
    return false
}
