package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type ItemGetEffectiveDeviceEnrollmentConfigurationsResponse struct {
    ItemGetEffectiveDeviceEnrollmentConfigurationsGetResponse
}
// NewItemGetEffectiveDeviceEnrollmentConfigurationsResponse instantiates a new ItemGetEffectiveDeviceEnrollmentConfigurationsResponse and sets the default values.
func NewItemGetEffectiveDeviceEnrollmentConfigurationsResponse()(*ItemGetEffectiveDeviceEnrollmentConfigurationsResponse) {
    m := &ItemGetEffectiveDeviceEnrollmentConfigurationsResponse{
        ItemGetEffectiveDeviceEnrollmentConfigurationsGetResponse: *NewItemGetEffectiveDeviceEnrollmentConfigurationsGetResponse(),
    }
    return m
}
// CreateItemGetEffectiveDeviceEnrollmentConfigurationsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemGetEffectiveDeviceEnrollmentConfigurationsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemGetEffectiveDeviceEnrollmentConfigurationsResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type ItemGetEffectiveDeviceEnrollmentConfigurationsResponseable interface {
    ItemGetEffectiveDeviceEnrollmentConfigurationsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
