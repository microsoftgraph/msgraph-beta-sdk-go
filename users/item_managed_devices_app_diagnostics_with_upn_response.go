package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemManagedDevicesAppDiagnosticsWithUpnGetResponseable instead.
type ItemManagedDevicesAppDiagnosticsWithUpnResponse struct {
    ItemManagedDevicesAppDiagnosticsWithUpnGetResponse
}
// NewItemManagedDevicesAppDiagnosticsWithUpnResponse instantiates a new ItemManagedDevicesAppDiagnosticsWithUpnResponse and sets the default values.
func NewItemManagedDevicesAppDiagnosticsWithUpnResponse()(*ItemManagedDevicesAppDiagnosticsWithUpnResponse) {
    m := &ItemManagedDevicesAppDiagnosticsWithUpnResponse{
        ItemManagedDevicesAppDiagnosticsWithUpnGetResponse: *NewItemManagedDevicesAppDiagnosticsWithUpnGetResponse(),
    }
    return m
}
// CreateItemManagedDevicesAppDiagnosticsWithUpnResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemManagedDevicesAppDiagnosticsWithUpnResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemManagedDevicesAppDiagnosticsWithUpnResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemManagedDevicesAppDiagnosticsWithUpnGetResponseable instead.
type ItemManagedDevicesAppDiagnosticsWithUpnResponseable interface {
    ItemManagedDevicesAppDiagnosticsWithUpnGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
