package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use OperationApprovalPoliciesRetrieveApprovableOperationsGetResponseable instead.
type OperationApprovalPoliciesRetrieveApprovableOperationsResponse struct {
    OperationApprovalPoliciesRetrieveApprovableOperationsGetResponse
}
// NewOperationApprovalPoliciesRetrieveApprovableOperationsResponse instantiates a new OperationApprovalPoliciesRetrieveApprovableOperationsResponse and sets the default values.
func NewOperationApprovalPoliciesRetrieveApprovableOperationsResponse()(*OperationApprovalPoliciesRetrieveApprovableOperationsResponse) {
    m := &OperationApprovalPoliciesRetrieveApprovableOperationsResponse{
        OperationApprovalPoliciesRetrieveApprovableOperationsGetResponse: *NewOperationApprovalPoliciesRetrieveApprovableOperationsGetResponse(),
    }
    return m
}
// CreateOperationApprovalPoliciesRetrieveApprovableOperationsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOperationApprovalPoliciesRetrieveApprovableOperationsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOperationApprovalPoliciesRetrieveApprovableOperationsResponse(), nil
}
// Deprecated: This class is obsolete. Use OperationApprovalPoliciesRetrieveApprovableOperationsGetResponseable instead.
type OperationApprovalPoliciesRetrieveApprovableOperationsResponseable interface {
    OperationApprovalPoliciesRetrieveApprovableOperationsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
