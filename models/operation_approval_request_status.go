package models
import (
    "errors"
)
// Indicates the status of the Approval Request. The status of a request will change when an action is successfully performed on it, such as when it is `approved` or `rejected`, or when the request's expiration DateTime passes and the result is `expired`.
type OperationApprovalRequestStatus int

const (
    // Default. Indicates that the request approval status is unknown, and that the status has not been assigned to the approval request.
    UNKNOWN_OPERATIONAPPROVALREQUESTSTATUS OperationApprovalRequestStatus = iota
    // Indicates that the approval request needs approval before the operation can be completed.
    NEEDSAPPROVAL_OPERATIONAPPROVALREQUESTSTATUS
    // Indicates that the approval request has been approved. The operation can now be completed.
    APPROVED_OPERATIONAPPROVALREQUESTSTATUS
    // Indicates that the approval request has been rejected. No further action can be taken to complete the operation or to modify the request.
    REJECTED_OPERATIONAPPROVALREQUESTSTATUS
    // Indicates that the approval request has been cancelled by the request's requestor. No further action can be taken to complete the operation or to modify the request.
    CANCELLED_OPERATIONAPPROVALREQUESTSTATUS
    // Indicates that the approval request has been completed. This status is feature agnostic and does not indicate success or failure of the operation. No further action is necessary for the operation or to modify the request.
    COMPLETED_OPERATIONAPPROVALREQUESTSTATUS
    // Indicates that the approval request has expired. No further action can be taken to complete the operation or to modify the request. A new approval request must be made and approved in order to complete the operation.
    EXPIRED_OPERATIONAPPROVALREQUESTSTATUS
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_OPERATIONAPPROVALREQUESTSTATUS
)

func (i OperationApprovalRequestStatus) String() string {
    return []string{"unknown", "needsApproval", "approved", "rejected", "cancelled", "completed", "expired", "unknownFutureValue"}[i]
}
func ParseOperationApprovalRequestStatus(v string) (any, error) {
    result := UNKNOWN_OPERATIONAPPROVALREQUESTSTATUS
    switch v {
        case "unknown":
            result = UNKNOWN_OPERATIONAPPROVALREQUESTSTATUS
        case "needsApproval":
            result = NEEDSAPPROVAL_OPERATIONAPPROVALREQUESTSTATUS
        case "approved":
            result = APPROVED_OPERATIONAPPROVALREQUESTSTATUS
        case "rejected":
            result = REJECTED_OPERATIONAPPROVALREQUESTSTATUS
        case "cancelled":
            result = CANCELLED_OPERATIONAPPROVALREQUESTSTATUS
        case "completed":
            result = COMPLETED_OPERATIONAPPROVALREQUESTSTATUS
        case "expired":
            result = EXPIRED_OPERATIONAPPROVALREQUESTSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_OPERATIONAPPROVALREQUESTSTATUS
        default:
            return 0, errors.New("Unknown OperationApprovalRequestStatus value: " + v)
    }
    return &result, nil
}
func SerializeOperationApprovalRequestStatus(values []OperationApprovalRequestStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OperationApprovalRequestStatus) isMultiValue() bool {
    return false
}
