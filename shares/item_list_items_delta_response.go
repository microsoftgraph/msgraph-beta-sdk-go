package shares

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemListItemsDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemListItemsDeltaResponse struct {
    ItemListItemsDeltaGetResponse
}
// NewItemListItemsDeltaResponse instantiates a new ItemListItemsDeltaResponse and sets the default values.
func NewItemListItemsDeltaResponse()(*ItemListItemsDeltaResponse) {
    m := &ItemListItemsDeltaResponse{
        ItemListItemsDeltaGetResponse: *NewItemListItemsDeltaGetResponse(),
    }
    return m
}
// CreateItemListItemsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemListItemsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemListItemsDeltaResponse(), nil
}
// ItemListItemsDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemListItemsDeltaResponseable interface {
    ItemListItemsDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
