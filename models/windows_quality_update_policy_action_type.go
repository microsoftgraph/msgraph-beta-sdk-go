// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models
// An enum type to represent approval actions of single or list of quality update policies
type WindowsQualityUpdatePolicyActionType int

const (
    // Enum value to represent the approval actions for quality update
    APPROVE_WINDOWSQUALITYUPDATEPOLICYACTIONTYPE WindowsQualityUpdatePolicyActionType = iota
    // Enum value to represent the pause actions for quality update
    SUSPEND_WINDOWSQUALITYUPDATEPOLICYACTIONTYPE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_WINDOWSQUALITYUPDATEPOLICYACTIONTYPE
)

func (i WindowsQualityUpdatePolicyActionType) String() string {
    return []string{"approve", "suspend", "unknownFutureValue"}[i]
}
func ParseWindowsQualityUpdatePolicyActionType(v string) (any, error) {
    result := APPROVE_WINDOWSQUALITYUPDATEPOLICYACTIONTYPE
    switch v {
        case "approve":
            result = APPROVE_WINDOWSQUALITYUPDATEPOLICYACTIONTYPE
        case "suspend":
            result = SUSPEND_WINDOWSQUALITYUPDATEPOLICYACTIONTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_WINDOWSQUALITYUPDATEPOLICYACTIONTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWindowsQualityUpdatePolicyActionType(values []WindowsQualityUpdatePolicyActionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WindowsQualityUpdatePolicyActionType) isMultiValue() bool {
    return false
}
