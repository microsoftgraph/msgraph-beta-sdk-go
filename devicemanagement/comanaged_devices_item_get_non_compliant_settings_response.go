package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type ComanagedDevicesItemGetNonCompliantSettingsResponse struct {
    ComanagedDevicesItemGetNonCompliantSettingsGetResponse
}
// NewComanagedDevicesItemGetNonCompliantSettingsResponse instantiates a new ComanagedDevicesItemGetNonCompliantSettingsResponse and sets the default values.
func NewComanagedDevicesItemGetNonCompliantSettingsResponse()(*ComanagedDevicesItemGetNonCompliantSettingsResponse) {
    m := &ComanagedDevicesItemGetNonCompliantSettingsResponse{
        ComanagedDevicesItemGetNonCompliantSettingsGetResponse: *NewComanagedDevicesItemGetNonCompliantSettingsGetResponse(),
    }
    return m
}
// CreateComanagedDevicesItemGetNonCompliantSettingsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateComanagedDevicesItemGetNonCompliantSettingsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewComanagedDevicesItemGetNonCompliantSettingsResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type ComanagedDevicesItemGetNonCompliantSettingsResponseable interface {
    ComanagedDevicesItemGetNonCompliantSettingsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
