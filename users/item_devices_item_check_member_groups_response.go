package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemDevicesItemCheckMemberGroupsResponse 
// Deprecated: This class is obsolete. Use checkMemberGroupsPostResponse instead.
type ItemDevicesItemCheckMemberGroupsResponse struct {
    ItemDevicesItemCheckMemberGroupsPostResponse
}
// NewItemDevicesItemCheckMemberGroupsResponse instantiates a new ItemDevicesItemCheckMemberGroupsResponse and sets the default values.
func NewItemDevicesItemCheckMemberGroupsResponse()(*ItemDevicesItemCheckMemberGroupsResponse) {
    m := &ItemDevicesItemCheckMemberGroupsResponse{
        ItemDevicesItemCheckMemberGroupsPostResponse: *NewItemDevicesItemCheckMemberGroupsPostResponse(),
    }
    return m
}
// CreateItemDevicesItemCheckMemberGroupsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemDevicesItemCheckMemberGroupsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemDevicesItemCheckMemberGroupsResponse(), nil
}
// ItemDevicesItemCheckMemberGroupsResponseable 
// Deprecated: This class is obsolete. Use checkMemberGroupsPostResponse instead.
type ItemDevicesItemCheckMemberGroupsResponseable interface {
    ItemDevicesItemCheckMemberGroupsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
