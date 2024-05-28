package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ManageddevicesItemGetfilevaultkeyGetFileVaultKeyGetResponseable instead.
type ManageddevicesItemGetfilevaultkeyGetFileVaultKeyResponse struct {
    ManageddevicesItemGetfilevaultkeyGetFileVaultKeyGetResponse
}
// NewManageddevicesItemGetfilevaultkeyGetFileVaultKeyResponse instantiates a new ManageddevicesItemGetfilevaultkeyGetFileVaultKeyResponse and sets the default values.
func NewManageddevicesItemGetfilevaultkeyGetFileVaultKeyResponse()(*ManageddevicesItemGetfilevaultkeyGetFileVaultKeyResponse) {
    m := &ManageddevicesItemGetfilevaultkeyGetFileVaultKeyResponse{
        ManageddevicesItemGetfilevaultkeyGetFileVaultKeyGetResponse: *NewManageddevicesItemGetfilevaultkeyGetFileVaultKeyGetResponse(),
    }
    return m
}
// CreateManageddevicesItemGetfilevaultkeyGetFileVaultKeyResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateManageddevicesItemGetfilevaultkeyGetFileVaultKeyResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManageddevicesItemGetfilevaultkeyGetFileVaultKeyResponse(), nil
}
// Deprecated: This class is obsolete. Use ManageddevicesItemGetfilevaultkeyGetFileVaultKeyGetResponseable instead.
type ManageddevicesItemGetfilevaultkeyGetFileVaultKeyResponseable interface {
    ManageddevicesItemGetfilevaultkeyGetFileVaultKeyGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
