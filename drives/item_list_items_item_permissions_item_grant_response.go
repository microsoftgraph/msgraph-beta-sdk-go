package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemListItemsItemPermissionsItemGrantPostResponseable instead.
type ItemListItemsItemPermissionsItemGrantResponse struct {
    ItemListItemsItemPermissionsItemGrantPostResponse
}
// NewItemListItemsItemPermissionsItemGrantResponse instantiates a new ItemListItemsItemPermissionsItemGrantResponse and sets the default values.
func NewItemListItemsItemPermissionsItemGrantResponse()(*ItemListItemsItemPermissionsItemGrantResponse) {
    m := &ItemListItemsItemPermissionsItemGrantResponse{
        ItemListItemsItemPermissionsItemGrantPostResponse: *NewItemListItemsItemPermissionsItemGrantPostResponse(),
    }
    return m
}
// CreateItemListItemsItemPermissionsItemGrantResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemListItemsItemPermissionsItemGrantResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemListItemsItemPermissionsItemGrantResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemListItemsItemPermissionsItemGrantPostResponseable instead.
type ItemListItemsItemPermissionsItemGrantResponseable interface {
    ItemListItemsItemPermissionsItemGrantPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
