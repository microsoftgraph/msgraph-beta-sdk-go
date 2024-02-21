package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use DeviceManagementScriptsHasPayloadLinksPostResponseable instead.
type DeviceManagementScriptsHasPayloadLinksResponse struct {
    DeviceManagementScriptsHasPayloadLinksPostResponse
}
// NewDeviceManagementScriptsHasPayloadLinksResponse instantiates a new DeviceManagementScriptsHasPayloadLinksResponse and sets the default values.
func NewDeviceManagementScriptsHasPayloadLinksResponse()(*DeviceManagementScriptsHasPayloadLinksResponse) {
    m := &DeviceManagementScriptsHasPayloadLinksResponse{
        DeviceManagementScriptsHasPayloadLinksPostResponse: *NewDeviceManagementScriptsHasPayloadLinksPostResponse(),
    }
    return m
}
// CreateDeviceManagementScriptsHasPayloadLinksResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceManagementScriptsHasPayloadLinksResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementScriptsHasPayloadLinksResponse(), nil
}
// Deprecated: This class is obsolete. Use DeviceManagementScriptsHasPayloadLinksPostResponseable instead.
type DeviceManagementScriptsHasPayloadLinksResponseable interface {
    DeviceManagementScriptsHasPayloadLinksPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
