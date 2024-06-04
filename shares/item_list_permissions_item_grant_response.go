package shares

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemListPermissionsItemGrantPostResponseable instead.
type ItemListPermissionsItemGrantResponse struct {
    ItemListPermissionsItemGrantPostResponse
}
// NewItemListPermissionsItemGrantResponse instantiates a new ItemListPermissionsItemGrantResponse and sets the default values.
func NewItemListPermissionsItemGrantResponse()(*ItemListPermissionsItemGrantResponse) {
    m := &ItemListPermissionsItemGrantResponse{
        ItemListPermissionsItemGrantPostResponse: *NewItemListPermissionsItemGrantPostResponse(),
    }
    return m
}
// CreateItemListPermissionsItemGrantResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemListPermissionsItemGrantResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemListPermissionsItemGrantResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemListPermissionsItemGrantPostResponseable instead.
type ItemListPermissionsItemGrantResponseable interface {
    ItemListPermissionsItemGrantPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
