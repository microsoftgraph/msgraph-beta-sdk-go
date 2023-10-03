package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ComanagedDevicesItemGetFileVaultKeyResponse 
// Deprecated: This class is obsolete. Use getFileVaultKeyGetResponse instead.
type ComanagedDevicesItemGetFileVaultKeyResponse struct {
    ComanagedDevicesItemGetFileVaultKeyGetResponse
}
// NewComanagedDevicesItemGetFileVaultKeyResponse instantiates a new ComanagedDevicesItemGetFileVaultKeyResponse and sets the default values.
func NewComanagedDevicesItemGetFileVaultKeyResponse()(*ComanagedDevicesItemGetFileVaultKeyResponse) {
    m := &ComanagedDevicesItemGetFileVaultKeyResponse{
        ComanagedDevicesItemGetFileVaultKeyGetResponse: *NewComanagedDevicesItemGetFileVaultKeyGetResponse(),
    }
    return m
}
// CreateComanagedDevicesItemGetFileVaultKeyResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateComanagedDevicesItemGetFileVaultKeyResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewComanagedDevicesItemGetFileVaultKeyResponse(), nil
}
// ComanagedDevicesItemGetFileVaultKeyResponseable 
// Deprecated: This class is obsolete. Use getFileVaultKeyGetResponse instead.
type ComanagedDevicesItemGetFileVaultKeyResponseable interface {
    ComanagedDevicesItemGetFileVaultKeyGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
