package graph
import (
    "strings"
    "errors"
)
// 
type ApprovalState int

const (
    PENDING_APPROVALSTATE ApprovalState = iota
    APPROVED_APPROVALSTATE
    DENIED_APPROVALSTATE
    ABORTED_APPROVALSTATE
    CANCELED_APPROVALSTATE
)

func (i ApprovalState) String() string {
    return []string{"PENDING", "APPROVED", "DENIED", "ABORTED", "CANCELED"}[i]
}
func ParseApprovalState(v string) (interface{}, error) {
    result := PENDING_APPROVALSTATE
    switch strings.ToUpper(v) {
        case "PENDING":
            result = PENDING_APPROVALSTATE
        case "APPROVED":
            result = APPROVED_APPROVALSTATE
        case "DENIED":
            result = DENIED_APPROVALSTATE
        case "ABORTED":
            result = ABORTED_APPROVALSTATE
        case "CANCELED":
            result = CANCELED_APPROVALSTATE
        default:
            return 0, errors.New("Unknown ApprovalState value: " + v)
    }
    return &result, nil
}
func SerializeApprovalState(values []ApprovalState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
