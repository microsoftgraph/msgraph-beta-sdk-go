package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type OperationApprovalPoliciesGetOperationsRequiringApprovalResponse struct {
    OperationApprovalPoliciesGetOperationsRequiringApprovalGetResponse
}
// NewOperationApprovalPoliciesGetOperationsRequiringApprovalResponse instantiates a new OperationApprovalPoliciesGetOperationsRequiringApprovalResponse and sets the default values.
func NewOperationApprovalPoliciesGetOperationsRequiringApprovalResponse()(*OperationApprovalPoliciesGetOperationsRequiringApprovalResponse) {
    m := &OperationApprovalPoliciesGetOperationsRequiringApprovalResponse{
        OperationApprovalPoliciesGetOperationsRequiringApprovalGetResponse: *NewOperationApprovalPoliciesGetOperationsRequiringApprovalGetResponse(),
    }
    return m
}
// CreateOperationApprovalPoliciesGetOperationsRequiringApprovalResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOperationApprovalPoliciesGetOperationsRequiringApprovalResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOperationApprovalPoliciesGetOperationsRequiringApprovalResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type OperationApprovalPoliciesGetOperationsRequiringApprovalResponseable interface {
    OperationApprovalPoliciesGetOperationsRequiringApprovalGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
