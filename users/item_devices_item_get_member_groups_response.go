package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemDevicesItemGetMemberGroupsResponse 
// Deprecated: This class is obsolete. Use getMemberGroupsPostResponse instead.
type ItemDevicesItemGetMemberGroupsResponse struct {
    ItemDevicesItemGetMemberGroupsPostResponse
}
// NewItemDevicesItemGetMemberGroupsResponse instantiates a new ItemDevicesItemGetMemberGroupsResponse and sets the default values.
func NewItemDevicesItemGetMemberGroupsResponse()(*ItemDevicesItemGetMemberGroupsResponse) {
    m := &ItemDevicesItemGetMemberGroupsResponse{
        ItemDevicesItemGetMemberGroupsPostResponse: *NewItemDevicesItemGetMemberGroupsPostResponse(),
    }
    return m
}
// CreateItemDevicesItemGetMemberGroupsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemDevicesItemGetMemberGroupsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemDevicesItemGetMemberGroupsResponse(), nil
}
// ItemDevicesItemGetMemberGroupsResponseable 
// Deprecated: This class is obsolete. Use getMemberGroupsPostResponse instead.
type ItemDevicesItemGetMemberGroupsResponseable interface {
    ItemDevicesItemGetMemberGroupsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
