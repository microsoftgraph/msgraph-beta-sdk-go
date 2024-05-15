package models
// The set of available policy types that can be configured for approval. The policy type must always be defined in an OperationApprovalRequest.
type OperationApprovalPolicyType int

const (
    // Default. Indicates that the configured policy type is unknown. Not a valid policy type in an OperationApprovalPolicy.
    UNKNOWN_OPERATIONAPPROVALPOLICYTYPE OperationApprovalPolicyType = iota
    // Indicates that the configured policy type is an application type, such as mobile apps or built-in apps.
    APP_OPERATIONAPPROVALPOLICYTYPE
    // Indicates that the configured policy type is a script type, such as Powershell scripts or remediation scripts.
    SCRIPT_OPERATIONAPPROVALPOLICYTYPE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_OPERATIONAPPROVALPOLICYTYPE
)

func (i OperationApprovalPolicyType) String() string {
    return []string{"unknown", "app", "script", "unknownFutureValue"}[i]
}
func ParseOperationApprovalPolicyType(v string) (any, error) {
    result := UNKNOWN_OPERATIONAPPROVALPOLICYTYPE
    switch v {
        case "unknown":
            result = UNKNOWN_OPERATIONAPPROVALPOLICYTYPE
        case "app":
            result = APP_OPERATIONAPPROVALPOLICYTYPE
        case "script":
            result = SCRIPT_OPERATIONAPPROVALPOLICYTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_OPERATIONAPPROVALPOLICYTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeOperationApprovalPolicyType(values []OperationApprovalPolicyType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OperationApprovalPolicyType) isMultiValue() bool {
    return false
}
