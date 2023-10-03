package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementScriptsHasPayloadLinksResponse 
// Deprecated: This class is obsolete. Use hasPayloadLinksPostResponse instead.
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
func CreateDeviceManagementScriptsHasPayloadLinksResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementScriptsHasPayloadLinksResponse(), nil
}
// DeviceManagementScriptsHasPayloadLinksResponseable 
// Deprecated: This class is obsolete. Use hasPayloadLinksPostResponse instead.
type DeviceManagementScriptsHasPayloadLinksResponseable interface {
    DeviceManagementScriptsHasPayloadLinksPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
