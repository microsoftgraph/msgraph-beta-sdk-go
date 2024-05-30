package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use DevicemanagementscriptsHaspayloadlinksHasPayloadLinksPostResponseable instead.
type DevicemanagementscriptsHaspayloadlinksHasPayloadLinksResponse struct {
    DevicemanagementscriptsHaspayloadlinksHasPayloadLinksPostResponse
}
// NewDevicemanagementscriptsHaspayloadlinksHasPayloadLinksResponse instantiates a new DevicemanagementscriptsHaspayloadlinksHasPayloadLinksResponse and sets the default values.
func NewDevicemanagementscriptsHaspayloadlinksHasPayloadLinksResponse()(*DevicemanagementscriptsHaspayloadlinksHasPayloadLinksResponse) {
    m := &DevicemanagementscriptsHaspayloadlinksHasPayloadLinksResponse{
        DevicemanagementscriptsHaspayloadlinksHasPayloadLinksPostResponse: *NewDevicemanagementscriptsHaspayloadlinksHasPayloadLinksPostResponse(),
    }
    return m
}
// CreateDevicemanagementscriptsHaspayloadlinksHasPayloadLinksResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDevicemanagementscriptsHaspayloadlinksHasPayloadLinksResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDevicemanagementscriptsHaspayloadlinksHasPayloadLinksResponse(), nil
}
// Deprecated: This class is obsolete. Use DevicemanagementscriptsHaspayloadlinksHasPayloadLinksPostResponseable instead.
type DevicemanagementscriptsHaspayloadlinksHasPayloadLinksResponseable interface {
    DevicemanagementscriptsHaspayloadlinksHasPayloadLinksPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
