package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemSitesItemListsItemPermissionsItemGrantPostResponseable instead.
type ItemSitesItemListsItemPermissionsItemGrantResponse struct {
    ItemSitesItemListsItemPermissionsItemGrantPostResponse
}
// NewItemSitesItemListsItemPermissionsItemGrantResponse instantiates a new ItemSitesItemListsItemPermissionsItemGrantResponse and sets the default values.
func NewItemSitesItemListsItemPermissionsItemGrantResponse()(*ItemSitesItemListsItemPermissionsItemGrantResponse) {
    m := &ItemSitesItemListsItemPermissionsItemGrantResponse{
        ItemSitesItemListsItemPermissionsItemGrantPostResponse: *NewItemSitesItemListsItemPermissionsItemGrantPostResponse(),
    }
    return m
}
// CreateItemSitesItemListsItemPermissionsItemGrantResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemSitesItemListsItemPermissionsItemGrantResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSitesItemListsItemPermissionsItemGrantResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemSitesItemListsItemPermissionsItemGrantPostResponseable instead.
type ItemSitesItemListsItemPermissionsItemGrantResponseable interface {
    ItemSitesItemListsItemPermissionsItemGrantPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
