package sites

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemListsItemItemsItemPermissionsItemGrantPostResponseable instead.
type ItemListsItemItemsItemPermissionsItemGrantResponse struct {
    ItemListsItemItemsItemPermissionsItemGrantPostResponse
}
// NewItemListsItemItemsItemPermissionsItemGrantResponse instantiates a new ItemListsItemItemsItemPermissionsItemGrantResponse and sets the default values.
func NewItemListsItemItemsItemPermissionsItemGrantResponse()(*ItemListsItemItemsItemPermissionsItemGrantResponse) {
    m := &ItemListsItemItemsItemPermissionsItemGrantResponse{
        ItemListsItemItemsItemPermissionsItemGrantPostResponse: *NewItemListsItemItemsItemPermissionsItemGrantPostResponse(),
    }
    return m
}
// CreateItemListsItemItemsItemPermissionsItemGrantResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemListsItemItemsItemPermissionsItemGrantResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemListsItemItemsItemPermissionsItemGrantResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemListsItemItemsItemPermissionsItemGrantPostResponseable instead.
type ItemListsItemItemsItemPermissionsItemGrantResponseable interface {
    ItemListsItemItemsItemPermissionsItemGrantPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
