package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use OperationApprovalPoliciesGetApprovableOperationsGetResponseable instead.
type OperationApprovalPoliciesGetApprovableOperationsResponse struct {
    OperationApprovalPoliciesGetApprovableOperationsGetResponse
}
// NewOperationApprovalPoliciesGetApprovableOperationsResponse instantiates a new OperationApprovalPoliciesGetApprovableOperationsResponse and sets the default values.
func NewOperationApprovalPoliciesGetApprovableOperationsResponse()(*OperationApprovalPoliciesGetApprovableOperationsResponse) {
    m := &OperationApprovalPoliciesGetApprovableOperationsResponse{
        OperationApprovalPoliciesGetApprovableOperationsGetResponse: *NewOperationApprovalPoliciesGetApprovableOperationsGetResponse(),
    }
    return m
}
// CreateOperationApprovalPoliciesGetApprovableOperationsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOperationApprovalPoliciesGetApprovableOperationsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOperationApprovalPoliciesGetApprovableOperationsResponse(), nil
}
// Deprecated: This class is obsolete. Use OperationApprovalPoliciesGetApprovableOperationsGetResponseable instead.
type OperationApprovalPoliciesGetApprovableOperationsResponseable interface {
    OperationApprovalPoliciesGetApprovableOperationsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
