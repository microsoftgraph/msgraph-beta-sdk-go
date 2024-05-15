package models
// Source of the action on the approval request.
type OperationApprovalSource int

const (
    // Default. Indicates the source of the action on the approval request is unknown.
    UNKNOWN_OPERATIONAPPROVALSOURCE OperationApprovalSource = iota
    // Indicates the source of the action on the approval request is from an interactive session using the Intune Admin Console.
    ADMINCONSOLE_OPERATIONAPPROVALSOURCE
    // Indicates the source of the action on the approval request is from an email based form.
    EMAIL_OPERATIONAPPROVALSOURCE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_OPERATIONAPPROVALSOURCE
)

func (i OperationApprovalSource) String() string {
    return []string{"unknown", "adminConsole", "email", "unknownFutureValue"}[i]
}
func ParseOperationApprovalSource(v string) (any, error) {
    result := UNKNOWN_OPERATIONAPPROVALSOURCE
    switch v {
        case "unknown":
            result = UNKNOWN_OPERATIONAPPROVALSOURCE
        case "adminConsole":
            result = ADMINCONSOLE_OPERATIONAPPROVALSOURCE
        case "email":
            result = EMAIL_OPERATIONAPPROVALSOURCE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_OPERATIONAPPROVALSOURCE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeOperationApprovalSource(values []OperationApprovalSource) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OperationApprovalSource) isMultiValue() bool {
    return false
}
