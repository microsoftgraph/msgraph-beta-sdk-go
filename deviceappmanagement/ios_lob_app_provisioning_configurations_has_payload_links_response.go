package deviceappmanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosLobAppProvisioningConfigurationsHasPayloadLinksResponse 
// Deprecated: This class is obsolete. Use hasPayloadLinksPostResponse instead.
type IosLobAppProvisioningConfigurationsHasPayloadLinksResponse struct {
    IosLobAppProvisioningConfigurationsHasPayloadLinksPostResponse
}
// NewIosLobAppProvisioningConfigurationsHasPayloadLinksResponse instantiates a new IosLobAppProvisioningConfigurationsHasPayloadLinksResponse and sets the default values.
func NewIosLobAppProvisioningConfigurationsHasPayloadLinksResponse()(*IosLobAppProvisioningConfigurationsHasPayloadLinksResponse) {
    m := &IosLobAppProvisioningConfigurationsHasPayloadLinksResponse{
        IosLobAppProvisioningConfigurationsHasPayloadLinksPostResponse: *NewIosLobAppProvisioningConfigurationsHasPayloadLinksPostResponse(),
    }
    return m
}
// CreateIosLobAppProvisioningConfigurationsHasPayloadLinksResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosLobAppProvisioningConfigurationsHasPayloadLinksResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosLobAppProvisioningConfigurationsHasPayloadLinksResponse(), nil
}
// IosLobAppProvisioningConfigurationsHasPayloadLinksResponseable 
// Deprecated: This class is obsolete. Use hasPayloadLinksPostResponse instead.
type IosLobAppProvisioningConfigurationsHasPayloadLinksResponseable interface {
    IosLobAppProvisioningConfigurationsHasPayloadLinksPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
