package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use DeviceConfigurationsGetIosAvailableUpdateVersionsGetResponseable instead.
type DeviceConfigurationsGetIosAvailableUpdateVersionsResponse struct {
    DeviceConfigurationsGetIosAvailableUpdateVersionsGetResponse
}
// NewDeviceConfigurationsGetIosAvailableUpdateVersionsResponse instantiates a new DeviceConfigurationsGetIosAvailableUpdateVersionsResponse and sets the default values.
func NewDeviceConfigurationsGetIosAvailableUpdateVersionsResponse()(*DeviceConfigurationsGetIosAvailableUpdateVersionsResponse) {
    m := &DeviceConfigurationsGetIosAvailableUpdateVersionsResponse{
        DeviceConfigurationsGetIosAvailableUpdateVersionsGetResponse: *NewDeviceConfigurationsGetIosAvailableUpdateVersionsGetResponse(),
    }
    return m
}
// CreateDeviceConfigurationsGetIosAvailableUpdateVersionsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceConfigurationsGetIosAvailableUpdateVersionsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceConfigurationsGetIosAvailableUpdateVersionsResponse(), nil
}
// Deprecated: This class is obsolete. Use DeviceConfigurationsGetIosAvailableUpdateVersionsGetResponseable instead.
type DeviceConfigurationsGetIosAvailableUpdateVersionsResponseable interface {
    DeviceConfigurationsGetIosAvailableUpdateVersionsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
