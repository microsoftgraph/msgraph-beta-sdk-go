package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceCompliancePoliciesHasPayloadLinksResponse 
// Deprecated: This class is obsolete. Use hasPayloadLinksPostResponse instead.
type DeviceCompliancePoliciesHasPayloadLinksResponse struct {
    DeviceCompliancePoliciesHasPayloadLinksPostResponse
}
// NewDeviceCompliancePoliciesHasPayloadLinksResponse instantiates a new DeviceCompliancePoliciesHasPayloadLinksResponse and sets the default values.
func NewDeviceCompliancePoliciesHasPayloadLinksResponse()(*DeviceCompliancePoliciesHasPayloadLinksResponse) {
    m := &DeviceCompliancePoliciesHasPayloadLinksResponse{
        DeviceCompliancePoliciesHasPayloadLinksPostResponse: *NewDeviceCompliancePoliciesHasPayloadLinksPostResponse(),
    }
    return m
}
// CreateDeviceCompliancePoliciesHasPayloadLinksResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceCompliancePoliciesHasPayloadLinksResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceCompliancePoliciesHasPayloadLinksResponse(), nil
}
// DeviceCompliancePoliciesHasPayloadLinksResponseable 
// Deprecated: This class is obsolete. Use hasPayloadLinksPostResponse instead.
type DeviceCompliancePoliciesHasPayloadLinksResponseable interface {
    DeviceCompliancePoliciesHasPayloadLinksPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
