package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceConfigurationsGetTargetedUsersAndDevicesResponse 
// Deprecated: This class is obsolete. Use getTargetedUsersAndDevicesPostResponse instead.
type DeviceConfigurationsGetTargetedUsersAndDevicesResponse struct {
    DeviceConfigurationsGetTargetedUsersAndDevicesPostResponse
}
// NewDeviceConfigurationsGetTargetedUsersAndDevicesResponse instantiates a new DeviceConfigurationsGetTargetedUsersAndDevicesResponse and sets the default values.
func NewDeviceConfigurationsGetTargetedUsersAndDevicesResponse()(*DeviceConfigurationsGetTargetedUsersAndDevicesResponse) {
    m := &DeviceConfigurationsGetTargetedUsersAndDevicesResponse{
        DeviceConfigurationsGetTargetedUsersAndDevicesPostResponse: *NewDeviceConfigurationsGetTargetedUsersAndDevicesPostResponse(),
    }
    return m
}
// CreateDeviceConfigurationsGetTargetedUsersAndDevicesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceConfigurationsGetTargetedUsersAndDevicesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceConfigurationsGetTargetedUsersAndDevicesResponse(), nil
}
// DeviceConfigurationsGetTargetedUsersAndDevicesResponseable 
// Deprecated: This class is obsolete. Use getTargetedUsersAndDevicesPostResponse instead.
type DeviceConfigurationsGetTargetedUsersAndDevicesResponseable interface {
    DeviceConfigurationsGetTargetedUsersAndDevicesPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
