package teamtemplatedefinition

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemTeamDefinitionPrimaryChannelMembersRemovePostResponseable instead.
type ItemTeamDefinitionPrimaryChannelMembersRemoveResponse struct {
    ItemTeamDefinitionPrimaryChannelMembersRemovePostResponse
}
// NewItemTeamDefinitionPrimaryChannelMembersRemoveResponse instantiates a new ItemTeamDefinitionPrimaryChannelMembersRemoveResponse and sets the default values.
func NewItemTeamDefinitionPrimaryChannelMembersRemoveResponse()(*ItemTeamDefinitionPrimaryChannelMembersRemoveResponse) {
    m := &ItemTeamDefinitionPrimaryChannelMembersRemoveResponse{
        ItemTeamDefinitionPrimaryChannelMembersRemovePostResponse: *NewItemTeamDefinitionPrimaryChannelMembersRemovePostResponse(),
    }
    return m
}
// CreateItemTeamDefinitionPrimaryChannelMembersRemoveResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamDefinitionPrimaryChannelMembersRemoveResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamDefinitionPrimaryChannelMembersRemoveResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemTeamDefinitionPrimaryChannelMembersRemovePostResponseable instead.
type ItemTeamDefinitionPrimaryChannelMembersRemoveResponseable interface {
    ItemTeamDefinitionPrimaryChannelMembersRemovePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
