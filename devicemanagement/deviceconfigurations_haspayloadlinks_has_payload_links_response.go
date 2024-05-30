package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use DeviceconfigurationsHaspayloadlinksHasPayloadLinksPostResponseable instead.
type DeviceconfigurationsHaspayloadlinksHasPayloadLinksResponse struct {
    DeviceconfigurationsHaspayloadlinksHasPayloadLinksPostResponse
}
// NewDeviceconfigurationsHaspayloadlinksHasPayloadLinksResponse instantiates a new DeviceconfigurationsHaspayloadlinksHasPayloadLinksResponse and sets the default values.
func NewDeviceconfigurationsHaspayloadlinksHasPayloadLinksResponse()(*DeviceconfigurationsHaspayloadlinksHasPayloadLinksResponse) {
    m := &DeviceconfigurationsHaspayloadlinksHasPayloadLinksResponse{
        DeviceconfigurationsHaspayloadlinksHasPayloadLinksPostResponse: *NewDeviceconfigurationsHaspayloadlinksHasPayloadLinksPostResponse(),
    }
    return m
}
// CreateDeviceconfigurationsHaspayloadlinksHasPayloadLinksResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceconfigurationsHaspayloadlinksHasPayloadLinksResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceconfigurationsHaspayloadlinksHasPayloadLinksResponse(), nil
}
// Deprecated: This class is obsolete. Use DeviceconfigurationsHaspayloadlinksHasPayloadLinksPostResponseable instead.
type DeviceconfigurationsHaspayloadlinksHasPayloadLinksResponseable interface {
    DeviceconfigurationsHaspayloadlinksHasPayloadLinksPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
