package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OperationApprovalRequestsItemCancelApprovalResponse 
// Deprecated: This class is obsolete. Use cancelApprovalPostResponse instead.
type OperationApprovalRequestsItemCancelApprovalResponse struct {
    OperationApprovalRequestsItemCancelApprovalPostResponse
}
// NewOperationApprovalRequestsItemCancelApprovalResponse instantiates a new OperationApprovalRequestsItemCancelApprovalResponse and sets the default values.
func NewOperationApprovalRequestsItemCancelApprovalResponse()(*OperationApprovalRequestsItemCancelApprovalResponse) {
    m := &OperationApprovalRequestsItemCancelApprovalResponse{
        OperationApprovalRequestsItemCancelApprovalPostResponse: *NewOperationApprovalRequestsItemCancelApprovalPostResponse(),
    }
    return m
}
// CreateOperationApprovalRequestsItemCancelApprovalResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOperationApprovalRequestsItemCancelApprovalResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOperationApprovalRequestsItemCancelApprovalResponse(), nil
}
// OperationApprovalRequestsItemCancelApprovalResponseable 
// Deprecated: This class is obsolete. Use cancelApprovalPostResponse instead.
type OperationApprovalRequestsItemCancelApprovalResponseable interface {
    OperationApprovalRequestsItemCancelApprovalPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
