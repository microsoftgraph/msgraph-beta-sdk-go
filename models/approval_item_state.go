package models
type ApprovalItemState int

const (
    CANCELED_APPROVALITEMSTATE ApprovalItemState = iota
    CREATED_APPROVALITEMSTATE
    PENDING_APPROVALITEMSTATE
    COMPLETED_APPROVALITEMSTATE
    UNKNOWNFUTUREVALUE_APPROVALITEMSTATE
)

func (i ApprovalItemState) String() string {
    return []string{"canceled", "created", "pending", "completed", "unknownFutureValue"}[i]
}
func ParseApprovalItemState(v string) (any, error) {
    result := CANCELED_APPROVALITEMSTATE
    switch v {
        case "canceled":
            result = CANCELED_APPROVALITEMSTATE
        case "created":
            result = CREATED_APPROVALITEMSTATE
        case "pending":
            result = PENDING_APPROVALITEMSTATE
        case "completed":
            result = COMPLETED_APPROVALITEMSTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_APPROVALITEMSTATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeApprovalItemState(values []ApprovalItemState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ApprovalItemState) isMultiValue() bool {
    return false
}
