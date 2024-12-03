package teamtemplatedefinition

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemTeamDefinitionChannelsItemGetAllMembersAddPostResponseable instead.
type ItemTeamDefinitionChannelsItemGetAllMembersAddResponse struct {
    ItemTeamDefinitionChannelsItemGetAllMembersAddPostResponse
}
// NewItemTeamDefinitionChannelsItemGetAllMembersAddResponse instantiates a new ItemTeamDefinitionChannelsItemGetAllMembersAddResponse and sets the default values.
func NewItemTeamDefinitionChannelsItemGetAllMembersAddResponse()(*ItemTeamDefinitionChannelsItemGetAllMembersAddResponse) {
    m := &ItemTeamDefinitionChannelsItemGetAllMembersAddResponse{
        ItemTeamDefinitionChannelsItemGetAllMembersAddPostResponse: *NewItemTeamDefinitionChannelsItemGetAllMembersAddPostResponse(),
    }
    return m
}
// CreateItemTeamDefinitionChannelsItemGetAllMembersAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamDefinitionChannelsItemGetAllMembersAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamDefinitionChannelsItemGetAllMembersAddResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemTeamDefinitionChannelsItemGetAllMembersAddPostResponseable instead.
type ItemTeamDefinitionChannelsItemGetAllMembersAddResponseable interface {
    ItemTeamDefinitionChannelsItemGetAllMembersAddPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
