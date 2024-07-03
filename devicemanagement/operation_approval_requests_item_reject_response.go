package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use OperationApprovalRequestsItemRejectPostResponseable instead.
type OperationApprovalRequestsItemRejectResponse struct {
    OperationApprovalRequestsItemRejectPostResponse
}
// NewOperationApprovalRequestsItemRejectResponse instantiates a new OperationApprovalRequestsItemRejectResponse and sets the default values.
func NewOperationApprovalRequestsItemRejectResponse()(*OperationApprovalRequestsItemRejectResponse) {
    m := &OperationApprovalRequestsItemRejectResponse{
        OperationApprovalRequestsItemRejectPostResponse: *NewOperationApprovalRequestsItemRejectPostResponse(),
    }
    return m
}
// CreateOperationApprovalRequestsItemRejectResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOperationApprovalRequestsItemRejectResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOperationApprovalRequestsItemRejectResponse(), nil
}
// Deprecated: This class is obsolete. Use OperationApprovalRequestsItemRejectPostResponseable instead.
type OperationApprovalRequestsItemRejectResponseable interface {
    OperationApprovalRequestsItemRejectPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
