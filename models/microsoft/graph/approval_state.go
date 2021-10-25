package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "PENDING":
            return PENDING_APPROVALSTATE, nil
        case "APPROVED":
            return APPROVED_APPROVALSTATE, nil
        case "DENIED":
            return DENIED_APPROVALSTATE, nil
        case "ABORTED":
            return ABORTED_APPROVALSTATE, nil
        case "CANCELED":
            return CANCELED_APPROVALSTATE, nil
    }
    return 0, errors.New("Unknown ApprovalState value: " + v)
}
func SerializeApprovalState(values []ApprovalState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
