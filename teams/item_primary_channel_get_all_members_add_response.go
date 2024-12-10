package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemPrimaryChannelGetAllMembersAddPostResponseable instead.
type ItemPrimaryChannelGetAllMembersAddResponse struct {
    ItemPrimaryChannelGetAllMembersAddPostResponse
}
// NewItemPrimaryChannelGetAllMembersAddResponse instantiates a new ItemPrimaryChannelGetAllMembersAddResponse and sets the default values.
func NewItemPrimaryChannelGetAllMembersAddResponse()(*ItemPrimaryChannelGetAllMembersAddResponse) {
    m := &ItemPrimaryChannelGetAllMembersAddResponse{
        ItemPrimaryChannelGetAllMembersAddPostResponse: *NewItemPrimaryChannelGetAllMembersAddPostResponse(),
    }
    return m
}
// CreateItemPrimaryChannelGetAllMembersAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemPrimaryChannelGetAllMembersAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPrimaryChannelGetAllMembersAddResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemPrimaryChannelGetAllMembersAddPostResponseable instead.
type ItemPrimaryChannelGetAllMembersAddResponseable interface {
    ItemPrimaryChannelGetAllMembersAddPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
