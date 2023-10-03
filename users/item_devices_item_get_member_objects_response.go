package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemDevicesItemGetMemberObjectsResponse 
// Deprecated: This class is obsolete. Use getMemberObjectsPostResponse instead.
type ItemDevicesItemGetMemberObjectsResponse struct {
    ItemDevicesItemGetMemberObjectsPostResponse
}
// NewItemDevicesItemGetMemberObjectsResponse instantiates a new ItemDevicesItemGetMemberObjectsResponse and sets the default values.
func NewItemDevicesItemGetMemberObjectsResponse()(*ItemDevicesItemGetMemberObjectsResponse) {
    m := &ItemDevicesItemGetMemberObjectsResponse{
        ItemDevicesItemGetMemberObjectsPostResponse: *NewItemDevicesItemGetMemberObjectsPostResponse(),
    }
    return m
}
// CreateItemDevicesItemGetMemberObjectsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemDevicesItemGetMemberObjectsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemDevicesItemGetMemberObjectsResponse(), nil
}
// ItemDevicesItemGetMemberObjectsResponseable 
// Deprecated: This class is obsolete. Use getMemberObjectsPostResponse instead.
type ItemDevicesItemGetMemberObjectsResponseable interface {
    ItemDevicesItemGetMemberObjectsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
