package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ManagedDevicesItemGetNonCompliantSettingsGetResponseable instead.
type ManagedDevicesItemGetNonCompliantSettingsResponse struct {
    ManagedDevicesItemGetNonCompliantSettingsGetResponse
}
// NewManagedDevicesItemGetNonCompliantSettingsResponse instantiates a new ManagedDevicesItemGetNonCompliantSettingsResponse and sets the default values.
func NewManagedDevicesItemGetNonCompliantSettingsResponse()(*ManagedDevicesItemGetNonCompliantSettingsResponse) {
    m := &ManagedDevicesItemGetNonCompliantSettingsResponse{
        ManagedDevicesItemGetNonCompliantSettingsGetResponse: *NewManagedDevicesItemGetNonCompliantSettingsGetResponse(),
    }
    return m
}
// CreateManagedDevicesItemGetNonCompliantSettingsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateManagedDevicesItemGetNonCompliantSettingsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedDevicesItemGetNonCompliantSettingsResponse(), nil
}
// Deprecated: This class is obsolete. Use ManagedDevicesItemGetNonCompliantSettingsGetResponseable instead.
type ManagedDevicesItemGetNonCompliantSettingsResponseable interface {
    ManagedDevicesItemGetNonCompliantSettingsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
