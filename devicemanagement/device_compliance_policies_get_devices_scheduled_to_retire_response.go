package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use DeviceCompliancePoliciesGetDevicesScheduledToRetireGetResponseable instead.
type DeviceCompliancePoliciesGetDevicesScheduledToRetireResponse struct {
    DeviceCompliancePoliciesGetDevicesScheduledToRetireGetResponse
}
// NewDeviceCompliancePoliciesGetDevicesScheduledToRetireResponse instantiates a new DeviceCompliancePoliciesGetDevicesScheduledToRetireResponse and sets the default values.
func NewDeviceCompliancePoliciesGetDevicesScheduledToRetireResponse()(*DeviceCompliancePoliciesGetDevicesScheduledToRetireResponse) {
    m := &DeviceCompliancePoliciesGetDevicesScheduledToRetireResponse{
        DeviceCompliancePoliciesGetDevicesScheduledToRetireGetResponse: *NewDeviceCompliancePoliciesGetDevicesScheduledToRetireGetResponse(),
    }
    return m
}
// CreateDeviceCompliancePoliciesGetDevicesScheduledToRetireResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceCompliancePoliciesGetDevicesScheduledToRetireResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceCompliancePoliciesGetDevicesScheduledToRetireResponse(), nil
}
// Deprecated: This class is obsolete. Use DeviceCompliancePoliciesGetDevicesScheduledToRetireGetResponseable instead.
type DeviceCompliancePoliciesGetDevicesScheduledToRetireResponseable interface {
    DeviceCompliancePoliciesGetDevicesScheduledToRetireGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
