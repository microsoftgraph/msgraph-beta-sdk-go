package sites

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemListsItemPermissionsItemGrantPostResponseable instead.
type ItemListsItemPermissionsItemGrantResponse struct {
    ItemListsItemPermissionsItemGrantPostResponse
}
// NewItemListsItemPermissionsItemGrantResponse instantiates a new ItemListsItemPermissionsItemGrantResponse and sets the default values.
func NewItemListsItemPermissionsItemGrantResponse()(*ItemListsItemPermissionsItemGrantResponse) {
    m := &ItemListsItemPermissionsItemGrantResponse{
        ItemListsItemPermissionsItemGrantPostResponse: *NewItemListsItemPermissionsItemGrantPostResponse(),
    }
    return m
}
// CreateItemListsItemPermissionsItemGrantResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemListsItemPermissionsItemGrantResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemListsItemPermissionsItemGrantResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemListsItemPermissionsItemGrantPostResponseable instead.
type ItemListsItemPermissionsItemGrantResponseable interface {
    ItemListsItemPermissionsItemGrantPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
