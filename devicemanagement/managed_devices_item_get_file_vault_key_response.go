package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedDevicesItemGetFileVaultKeyResponse 
// Deprecated: This class is obsolete. Use getFileVaultKeyGetResponse instead.
type ManagedDevicesItemGetFileVaultKeyResponse struct {
    ManagedDevicesItemGetFileVaultKeyGetResponse
}
// NewManagedDevicesItemGetFileVaultKeyResponse instantiates a new ManagedDevicesItemGetFileVaultKeyResponse and sets the default values.
func NewManagedDevicesItemGetFileVaultKeyResponse()(*ManagedDevicesItemGetFileVaultKeyResponse) {
    m := &ManagedDevicesItemGetFileVaultKeyResponse{
        ManagedDevicesItemGetFileVaultKeyGetResponse: *NewManagedDevicesItemGetFileVaultKeyGetResponse(),
    }
    return m
}
// CreateManagedDevicesItemGetFileVaultKeyResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedDevicesItemGetFileVaultKeyResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedDevicesItemGetFileVaultKeyResponse(), nil
}
// ManagedDevicesItemGetFileVaultKeyResponseable 
// Deprecated: This class is obsolete. Use getFileVaultKeyGetResponse instead.
type ManagedDevicesItemGetFileVaultKeyResponseable interface {
    ManagedDevicesItemGetFileVaultKeyGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
