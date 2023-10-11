package models
import (
    "errors"
)
// Indicates the type of elevated process
type PrivilegeManagementProcessType int

const (
    // Default. If the type was unknown on the client for some reasons
    UNDEFINED_PRIVILEGEMANAGEMENTPROCESSTYPE PrivilegeManagementProcessType = iota
    // The elevated process is a parent process
    PARENT_PRIVILEGEMANAGEMENTPROCESSTYPE
    // The elevated process is a child process
    CHILD_PRIVILEGEMANAGEMENTPROCESSTYPE
    // Evolvable emuneration sentinel value. Do not use
    UNKNOWNFUTUREVALUE_PRIVILEGEMANAGEMENTPROCESSTYPE
)

func (i PrivilegeManagementProcessType) String() string {
    return []string{"undefined", "parent", "child", "unknownFutureValue"}[i]
}
func ParsePrivilegeManagementProcessType(v string) (any, error) {
    result := UNDEFINED_PRIVILEGEMANAGEMENTPROCESSTYPE
    switch v {
        case "undefined":
            result = UNDEFINED_PRIVILEGEMANAGEMENTPROCESSTYPE
        case "parent":
            result = PARENT_PRIVILEGEMANAGEMENTPROCESSTYPE
        case "child":
            result = CHILD_PRIVILEGEMANAGEMENTPROCESSTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PRIVILEGEMANAGEMENTPROCESSTYPE
        default:
            return 0, errors.New("Unknown PrivilegeManagementProcessType value: " + v)
    }
    return &result, nil
}
func SerializePrivilegeManagementProcessType(values []PrivilegeManagementProcessType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PrivilegeManagementProcessType) isMultiValue() bool {
    return false
}
