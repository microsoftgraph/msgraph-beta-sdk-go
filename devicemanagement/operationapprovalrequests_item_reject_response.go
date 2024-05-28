package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use OperationapprovalrequestsItemRejectPostResponseable instead.
type OperationapprovalrequestsItemRejectResponse struct {
    OperationapprovalrequestsItemRejectPostResponse
}
// NewOperationapprovalrequestsItemRejectResponse instantiates a new OperationapprovalrequestsItemRejectResponse and sets the default values.
func NewOperationapprovalrequestsItemRejectResponse()(*OperationapprovalrequestsItemRejectResponse) {
    m := &OperationapprovalrequestsItemRejectResponse{
        OperationapprovalrequestsItemRejectPostResponse: *NewOperationapprovalrequestsItemRejectPostResponse(),
    }
    return m
}
// CreateOperationapprovalrequestsItemRejectResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOperationapprovalrequestsItemRejectResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOperationapprovalrequestsItemRejectResponse(), nil
}
// Deprecated: This class is obsolete. Use OperationapprovalrequestsItemRejectPostResponseable instead.
type OperationapprovalrequestsItemRejectResponseable interface {
    OperationapprovalrequestsItemRejectPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
