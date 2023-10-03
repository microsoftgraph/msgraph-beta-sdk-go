package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemDevicesItemCheckMemberObjectsResponse 
// Deprecated: This class is obsolete. Use checkMemberObjectsPostResponse instead.
type ItemDevicesItemCheckMemberObjectsResponse struct {
    ItemDevicesItemCheckMemberObjectsPostResponse
}
// NewItemDevicesItemCheckMemberObjectsResponse instantiates a new ItemDevicesItemCheckMemberObjectsResponse and sets the default values.
func NewItemDevicesItemCheckMemberObjectsResponse()(*ItemDevicesItemCheckMemberObjectsResponse) {
    m := &ItemDevicesItemCheckMemberObjectsResponse{
        ItemDevicesItemCheckMemberObjectsPostResponse: *NewItemDevicesItemCheckMemberObjectsPostResponse(),
    }
    return m
}
// CreateItemDevicesItemCheckMemberObjectsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemDevicesItemCheckMemberObjectsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemDevicesItemCheckMemberObjectsResponse(), nil
}
// ItemDevicesItemCheckMemberObjectsResponseable 
// Deprecated: This class is obsolete. Use checkMemberObjectsPostResponse instead.
type ItemDevicesItemCheckMemberObjectsResponseable interface {
    ItemDevicesItemCheckMemberObjectsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
