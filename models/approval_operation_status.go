package models
type ApprovalOperationStatus int

const (
    SCHEDULED_APPROVALOPERATIONSTATUS ApprovalOperationStatus = iota
    INPROGRESS_APPROVALOPERATIONSTATUS
    SUCCEEDED_APPROVALOPERATIONSTATUS
    FAILED_APPROVALOPERATIONSTATUS
    TIMEOUT_APPROVALOPERATIONSTATUS
    UNKNOWNFUTUREVALUE_APPROVALOPERATIONSTATUS
)

func (i ApprovalOperationStatus) String() string {
    return []string{"scheduled", "inProgress", "succeeded", "failed", "timeout", "unknownFutureValue"}[i]
}
func ParseApprovalOperationStatus(v string) (any, error) {
    result := SCHEDULED_APPROVALOPERATIONSTATUS
    switch v {
        case "scheduled":
            result = SCHEDULED_APPROVALOPERATIONSTATUS
        case "inProgress":
            result = INPROGRESS_APPROVALOPERATIONSTATUS
        case "succeeded":
            result = SUCCEEDED_APPROVALOPERATIONSTATUS
        case "failed":
            result = FAILED_APPROVALOPERATIONSTATUS
        case "timeout":
            result = TIMEOUT_APPROVALOPERATIONSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_APPROVALOPERATIONSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeApprovalOperationStatus(values []ApprovalOperationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ApprovalOperationStatus) isMultiValue() bool {
    return false
}
