package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemChannelsItemGetAllMembersRemovePostResponseable instead.
type ItemChannelsItemGetAllMembersRemoveResponse struct {
    ItemChannelsItemGetAllMembersRemovePostResponse
}
// NewItemChannelsItemGetAllMembersRemoveResponse instantiates a new ItemChannelsItemGetAllMembersRemoveResponse and sets the default values.
func NewItemChannelsItemGetAllMembersRemoveResponse()(*ItemChannelsItemGetAllMembersRemoveResponse) {
    m := &ItemChannelsItemGetAllMembersRemoveResponse{
        ItemChannelsItemGetAllMembersRemovePostResponse: *NewItemChannelsItemGetAllMembersRemovePostResponse(),
    }
    return m
}
// CreateItemChannelsItemGetAllMembersRemoveResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemChannelsItemGetAllMembersRemoveResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemChannelsItemGetAllMembersRemoveResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemChannelsItemGetAllMembersRemovePostResponseable instead.
type ItemChannelsItemGetAllMembersRemoveResponseable interface {
    ItemChannelsItemGetAllMembersRemovePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
