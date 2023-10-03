package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceConfigurationsHasPayloadLinksResponse 
// Deprecated: This class is obsolete. Use hasPayloadLinksPostResponse instead.
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
func CreateDeviceConfigurationsHasPayloadLinksResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceConfigurationsHasPayloadLinksResponse(), nil
}
// DeviceConfigurationsHasPayloadLinksResponseable 
// Deprecated: This class is obsolete. Use hasPayloadLinksPostResponse instead.
type DeviceConfigurationsHasPayloadLinksResponseable interface {
    DeviceConfigurationsHasPayloadLinksPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
