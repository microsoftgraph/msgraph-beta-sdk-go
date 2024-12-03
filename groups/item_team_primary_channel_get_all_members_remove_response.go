package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemTeamPrimaryChannelGetAllMembersRemovePostResponseable instead.
type ItemTeamPrimaryChannelGetAllMembersRemoveResponse struct {
    ItemTeamPrimaryChannelGetAllMembersRemovePostResponse
}
// NewItemTeamPrimaryChannelGetAllMembersRemoveResponse instantiates a new ItemTeamPrimaryChannelGetAllMembersRemoveResponse and sets the default values.
func NewItemTeamPrimaryChannelGetAllMembersRemoveResponse()(*ItemTeamPrimaryChannelGetAllMembersRemoveResponse) {
    m := &ItemTeamPrimaryChannelGetAllMembersRemoveResponse{
        ItemTeamPrimaryChannelGetAllMembersRemovePostResponse: *NewItemTeamPrimaryChannelGetAllMembersRemovePostResponse(),
    }
    return m
}
// CreateItemTeamPrimaryChannelGetAllMembersRemoveResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamPrimaryChannelGetAllMembersRemoveResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamPrimaryChannelGetAllMembersRemoveResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemTeamPrimaryChannelGetAllMembersRemovePostResponseable instead.
type ItemTeamPrimaryChannelGetAllMembersRemoveResponseable interface {
    ItemTeamPrimaryChannelGetAllMembersRemovePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
