package sites

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemListsItemItemsDeltaWithTokenResponse 
// Deprecated: This class is obsolete. Use deltaWithTokenGetResponse instead.
type ItemListsItemItemsDeltaWithTokenResponse struct {
    ItemListsItemItemsDeltaWithTokenGetResponse
}
// NewItemListsItemItemsDeltaWithTokenResponse instantiates a new ItemListsItemItemsDeltaWithTokenResponse and sets the default values.
func NewItemListsItemItemsDeltaWithTokenResponse()(*ItemListsItemItemsDeltaWithTokenResponse) {
    m := &ItemListsItemItemsDeltaWithTokenResponse{
        ItemListsItemItemsDeltaWithTokenGetResponse: *NewItemListsItemItemsDeltaWithTokenGetResponse(),
    }
    return m
}
// CreateItemListsItemItemsDeltaWithTokenResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemListsItemItemsDeltaWithTokenResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemListsItemItemsDeltaWithTokenResponse(), nil
}
// ItemListsItemItemsDeltaWithTokenResponseable 
// Deprecated: This class is obsolete. Use deltaWithTokenGetResponse instead.
type ItemListsItemItemsDeltaWithTokenResponseable interface {
    ItemListsItemItemsDeltaWithTokenGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
