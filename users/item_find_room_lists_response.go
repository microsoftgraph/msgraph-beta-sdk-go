package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type ItemFindRoomListsResponse struct {
    ItemFindRoomListsGetResponse
}
// NewItemFindRoomListsResponse instantiates a new ItemFindRoomListsResponse and sets the default values.
func NewItemFindRoomListsResponse()(*ItemFindRoomListsResponse) {
    m := &ItemFindRoomListsResponse{
        ItemFindRoomListsGetResponse: *NewItemFindRoomListsGetResponse(),
    }
    return m
}
// CreateItemFindRoomListsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemFindRoomListsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemFindRoomListsResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type ItemFindRoomListsResponseable interface {
    ItemFindRoomListsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
