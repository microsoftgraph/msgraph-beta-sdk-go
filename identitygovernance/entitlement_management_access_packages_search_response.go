package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EntitlementManagementAccessPackagesSearchResponse 
// Deprecated: This class is obsolete. Use SearchGetResponse instead.
type EntitlementManagementAccessPackagesSearchResponse struct {
    EntitlementManagementAccessPackagesSearchGetResponse
}
// NewEntitlementManagementAccessPackagesSearchResponse instantiates a new EntitlementManagementAccessPackagesSearchResponse and sets the default values.
func NewEntitlementManagementAccessPackagesSearchResponse()(*EntitlementManagementAccessPackagesSearchResponse) {
    m := &EntitlementManagementAccessPackagesSearchResponse{
        EntitlementManagementAccessPackagesSearchGetResponse: *NewEntitlementManagementAccessPackagesSearchGetResponse(),
    }
    return m
}
// CreateEntitlementManagementAccessPackagesSearchResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEntitlementManagementAccessPackagesSearchResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEntitlementManagementAccessPackagesSearchResponse(), nil
}
// EntitlementManagementAccessPackagesSearchResponseable 
// Deprecated: This class is obsolete. Use SearchGetResponse instead.
type EntitlementManagementAccessPackagesSearchResponseable interface {
    EntitlementManagementAccessPackagesSearchGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
