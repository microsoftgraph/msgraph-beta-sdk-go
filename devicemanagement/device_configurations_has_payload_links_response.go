package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type DeviceConfigurationsHasPayloadLinksResponse struct {
    DeviceConfigurationsHasPayloadLinksPostResponse
}
// NewDeviceConfigurationsHasPayloadLinksResponse instantiates a new DeviceConfigurationsHasPayloadLinksResponse and sets the default values.
func NewDeviceConfigurationsHasPayloadLinksResponse()(*DeviceConfigurationsHasPayloadLinksResponse) {
    m := &DeviceConfigurationsHasPayloadLinksResponse{
        DeviceConfigurationsHasPayloadLinksPostResponse: *NewDeviceConfigurationsHasPayloadLinksPostResponse(),
    }
    return m
}
// CreateDeviceConfigurationsHasPayloadLinksResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceConfigurationsHasPayloadLinksResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceConfigurationsHasPayloadLinksResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type DeviceConfigurationsHasPayloadLinksResponseable interface {
    DeviceConfigurationsHasPayloadLinksPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
