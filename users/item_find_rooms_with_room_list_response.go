package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemFindRoomsWithRoomListGetResponseable instead.
type ItemFindRoomsWithRoomListResponse struct {
    ItemFindRoomsWithRoomListGetResponse
}
// NewItemFindRoomsWithRoomListResponse instantiates a new ItemFindRoomsWithRoomListResponse and sets the default values.
func NewItemFindRoomsWithRoomListResponse()(*ItemFindRoomsWithRoomListResponse) {
    m := &ItemFindRoomsWithRoomListResponse{
        ItemFindRoomsWithRoomListGetResponse: *NewItemFindRoomsWithRoomListGetResponse(),
    }
    return m
}
// CreateItemFindRoomsWithRoomListResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemFindRoomsWithRoomListResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemFindRoomsWithRoomListResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemFindRoomsWithRoomListGetResponseable instead.
type ItemFindRoomsWithRoomListResponseable interface {
    ItemFindRoomsWithRoomListGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
