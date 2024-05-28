package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ComanageddevicesItemGetfilevaultkeyGetFileVaultKeyGetResponseable instead.
type ComanageddevicesItemGetfilevaultkeyGetFileVaultKeyResponse struct {
    ComanageddevicesItemGetfilevaultkeyGetFileVaultKeyGetResponse
}
// NewComanageddevicesItemGetfilevaultkeyGetFileVaultKeyResponse instantiates a new ComanageddevicesItemGetfilevaultkeyGetFileVaultKeyResponse and sets the default values.
func NewComanageddevicesItemGetfilevaultkeyGetFileVaultKeyResponse()(*ComanageddevicesItemGetfilevaultkeyGetFileVaultKeyResponse) {
    m := &ComanageddevicesItemGetfilevaultkeyGetFileVaultKeyResponse{
        ComanageddevicesItemGetfilevaultkeyGetFileVaultKeyGetResponse: *NewComanageddevicesItemGetfilevaultkeyGetFileVaultKeyGetResponse(),
    }
    return m
}
// CreateComanageddevicesItemGetfilevaultkeyGetFileVaultKeyResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateComanageddevicesItemGetfilevaultkeyGetFileVaultKeyResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewComanageddevicesItemGetfilevaultkeyGetFileVaultKeyResponse(), nil
}
// Deprecated: This class is obsolete. Use ComanageddevicesItemGetfilevaultkeyGetFileVaultKeyGetResponseable instead.
type ComanageddevicesItemGetfilevaultkeyGetFileVaultKeyResponseable interface {
    ComanageddevicesItemGetfilevaultkeyGetFileVaultKeyGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
