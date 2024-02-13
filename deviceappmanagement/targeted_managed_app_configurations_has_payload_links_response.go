package deviceappmanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type TargetedManagedAppConfigurationsHasPayloadLinksResponse struct {
    TargetedManagedAppConfigurationsHasPayloadLinksPostResponse
}
// NewTargetedManagedAppConfigurationsHasPayloadLinksResponse instantiates a new TargetedManagedAppConfigurationsHasPayloadLinksResponse and sets the default values.
func NewTargetedManagedAppConfigurationsHasPayloadLinksResponse()(*TargetedManagedAppConfigurationsHasPayloadLinksResponse) {
    m := &TargetedManagedAppConfigurationsHasPayloadLinksResponse{
        TargetedManagedAppConfigurationsHasPayloadLinksPostResponse: *NewTargetedManagedAppConfigurationsHasPayloadLinksPostResponse(),
    }
    return m
}
// CreateTargetedManagedAppConfigurationsHasPayloadLinksResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTargetedManagedAppConfigurationsHasPayloadLinksResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTargetedManagedAppConfigurationsHasPayloadLinksResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type TargetedManagedAppConfigurationsHasPayloadLinksResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TargetedManagedAppConfigurationsHasPayloadLinksPostResponseable
}
