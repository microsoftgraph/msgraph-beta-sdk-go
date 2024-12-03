package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemPrimaryChannelGetAllMembersRemovePostResponseable instead.
type ItemPrimaryChannelGetAllMembersRemoveResponse struct {
    ItemPrimaryChannelGetAllMembersRemovePostResponse
}
// NewItemPrimaryChannelGetAllMembersRemoveResponse instantiates a new ItemPrimaryChannelGetAllMembersRemoveResponse and sets the default values.
func NewItemPrimaryChannelGetAllMembersRemoveResponse()(*ItemPrimaryChannelGetAllMembersRemoveResponse) {
    m := &ItemPrimaryChannelGetAllMembersRemoveResponse{
        ItemPrimaryChannelGetAllMembersRemovePostResponse: *NewItemPrimaryChannelGetAllMembersRemovePostResponse(),
    }
    return m
}
// CreateItemPrimaryChannelGetAllMembersRemoveResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemPrimaryChannelGetAllMembersRemoveResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPrimaryChannelGetAllMembersRemoveResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemPrimaryChannelGetAllMembersRemovePostResponseable instead.
type ItemPrimaryChannelGetAllMembersRemoveResponseable interface {
    ItemPrimaryChannelGetAllMembersRemovePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
