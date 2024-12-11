package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemTeamPrimaryChannelGetAllMembersAddPostResponseable instead.
type ItemTeamPrimaryChannelGetAllMembersAddResponse struct {
    ItemTeamPrimaryChannelGetAllMembersAddPostResponse
}
// NewItemTeamPrimaryChannelGetAllMembersAddResponse instantiates a new ItemTeamPrimaryChannelGetAllMembersAddResponse and sets the default values.
func NewItemTeamPrimaryChannelGetAllMembersAddResponse()(*ItemTeamPrimaryChannelGetAllMembersAddResponse) {
    m := &ItemTeamPrimaryChannelGetAllMembersAddResponse{
        ItemTeamPrimaryChannelGetAllMembersAddPostResponse: *NewItemTeamPrimaryChannelGetAllMembersAddPostResponse(),
    }
    return m
}
// CreateItemTeamPrimaryChannelGetAllMembersAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamPrimaryChannelGetAllMembersAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamPrimaryChannelGetAllMembersAddResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemTeamPrimaryChannelGetAllMembersAddPostResponseable instead.
type ItemTeamPrimaryChannelGetAllMembersAddResponseable interface {
    ItemTeamPrimaryChannelGetAllMembersAddPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
