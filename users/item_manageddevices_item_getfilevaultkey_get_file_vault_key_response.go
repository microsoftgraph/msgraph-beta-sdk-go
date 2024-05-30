package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyGetResponseable instead.
type ItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyResponse struct {
    ItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyGetResponse
}
// NewItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyResponse instantiates a new ItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyResponse and sets the default values.
func NewItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyResponse()(*ItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyResponse) {
    m := &ItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyResponse{
        ItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyGetResponse: *NewItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyGetResponse(),
    }
    return m
}
// CreateItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyGetResponseable instead.
type ItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyResponseable interface {
    ItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
