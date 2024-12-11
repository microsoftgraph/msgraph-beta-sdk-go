package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemTeamChannelsItemGetAllMembersAddPostResponseable instead.
type ItemTeamChannelsItemGetAllMembersAddResponse struct {
    ItemTeamChannelsItemGetAllMembersAddPostResponse
}
// NewItemTeamChannelsItemGetAllMembersAddResponse instantiates a new ItemTeamChannelsItemGetAllMembersAddResponse and sets the default values.
func NewItemTeamChannelsItemGetAllMembersAddResponse()(*ItemTeamChannelsItemGetAllMembersAddResponse) {
    m := &ItemTeamChannelsItemGetAllMembersAddResponse{
        ItemTeamChannelsItemGetAllMembersAddPostResponse: *NewItemTeamChannelsItemGetAllMembersAddPostResponse(),
    }
    return m
}
// CreateItemTeamChannelsItemGetAllMembersAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamChannelsItemGetAllMembersAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamChannelsItemGetAllMembersAddResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemTeamChannelsItemGetAllMembersAddPostResponseable instead.
type ItemTeamChannelsItemGetAllMembersAddResponseable interface {
    ItemTeamChannelsItemGetAllMembersAddPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
