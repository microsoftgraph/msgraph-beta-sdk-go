// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use EntitlementManagementAccessPackageAssignmentRequestsFilterByCurrentUserWithOnGetResponseable instead.
type EntitlementManagementAccessPackageAssignmentRequestsFilterByCurrentUserWithOnResponse struct {
    EntitlementManagementAccessPackageAssignmentRequestsFilterByCurrentUserWithOnGetResponse
}
// NewEntitlementManagementAccessPackageAssignmentRequestsFilterByCurrentUserWithOnResponse instantiates a new EntitlementManagementAccessPackageAssignmentRequestsFilterByCurrentUserWithOnResponse and sets the default values.
func NewEntitlementManagementAccessPackageAssignmentRequestsFilterByCurrentUserWithOnResponse()(*EntitlementManagementAccessPackageAssignmentRequestsFilterByCurrentUserWithOnResponse) {
    m := &EntitlementManagementAccessPackageAssignmentRequestsFilterByCurrentUserWithOnResponse{
        EntitlementManagementAccessPackageAssignmentRequestsFilterByCurrentUserWithOnGetResponse: *NewEntitlementManagementAccessPackageAssignmentRequestsFilterByCurrentUserWithOnGetResponse(),
    }
    return m
}
// CreateEntitlementManagementAccessPackageAssignmentRequestsFilterByCurrentUserWithOnResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEntitlementManagementAccessPackageAssignmentRequestsFilterByCurrentUserWithOnResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEntitlementManagementAccessPackageAssignmentRequestsFilterByCurrentUserWithOnResponse(), nil
}
// Deprecated: This class is obsolete. Use EntitlementManagementAccessPackageAssignmentRequestsFilterByCurrentUserWithOnGetResponseable instead.
type EntitlementManagementAccessPackageAssignmentRequestsFilterByCurrentUserWithOnResponseable interface {
    EntitlementManagementAccessPackageAssignmentRequestsFilterByCurrentUserWithOnGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
