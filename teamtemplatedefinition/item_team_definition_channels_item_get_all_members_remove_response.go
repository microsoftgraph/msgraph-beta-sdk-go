package teamtemplatedefinition

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemTeamDefinitionChannelsItemGetAllMembersRemovePostResponseable instead.
type ItemTeamDefinitionChannelsItemGetAllMembersRemoveResponse struct {
    ItemTeamDefinitionChannelsItemGetAllMembersRemovePostResponse
}
// NewItemTeamDefinitionChannelsItemGetAllMembersRemoveResponse instantiates a new ItemTeamDefinitionChannelsItemGetAllMembersRemoveResponse and sets the default values.
func NewItemTeamDefinitionChannelsItemGetAllMembersRemoveResponse()(*ItemTeamDefinitionChannelsItemGetAllMembersRemoveResponse) {
    m := &ItemTeamDefinitionChannelsItemGetAllMembersRemoveResponse{
        ItemTeamDefinitionChannelsItemGetAllMembersRemovePostResponse: *NewItemTeamDefinitionChannelsItemGetAllMembersRemovePostResponse(),
    }
    return m
}
// CreateItemTeamDefinitionChannelsItemGetAllMembersRemoveResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamDefinitionChannelsItemGetAllMembersRemoveResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamDefinitionChannelsItemGetAllMembersRemoveResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemTeamDefinitionChannelsItemGetAllMembersRemovePostResponseable instead.
type ItemTeamDefinitionChannelsItemGetAllMembersRemoveResponseable interface {
    ItemTeamDefinitionChannelsItemGetAllMembersRemovePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
