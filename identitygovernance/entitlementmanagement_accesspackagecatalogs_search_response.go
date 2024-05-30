package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use EntitlementmanagementAccesspackagecatalogsSearchGetResponseable instead.
type EntitlementmanagementAccesspackagecatalogsSearchResponse struct {
    EntitlementmanagementAccesspackagecatalogsSearchGetResponse
}
// NewEntitlementmanagementAccesspackagecatalogsSearchResponse instantiates a new EntitlementmanagementAccesspackagecatalogsSearchResponse and sets the default values.
func NewEntitlementmanagementAccesspackagecatalogsSearchResponse()(*EntitlementmanagementAccesspackagecatalogsSearchResponse) {
    m := &EntitlementmanagementAccesspackagecatalogsSearchResponse{
        EntitlementmanagementAccesspackagecatalogsSearchGetResponse: *NewEntitlementmanagementAccesspackagecatalogsSearchGetResponse(),
    }
    return m
}
// CreateEntitlementmanagementAccesspackagecatalogsSearchResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEntitlementmanagementAccesspackagecatalogsSearchResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEntitlementmanagementAccesspackagecatalogsSearchResponse(), nil
}
// Deprecated: This class is obsolete. Use EntitlementmanagementAccesspackagecatalogsSearchGetResponseable instead.
type EntitlementmanagementAccesspackagecatalogsSearchResponseable interface {
    EntitlementmanagementAccesspackagecatalogsSearchGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
