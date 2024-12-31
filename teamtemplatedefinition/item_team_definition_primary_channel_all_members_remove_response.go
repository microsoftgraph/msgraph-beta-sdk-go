package teamtemplatedefinition

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemTeamDefinitionPrimaryChannelAllMembersRemovePostResponseable instead.
type ItemTeamDefinitionPrimaryChannelAllMembersRemoveResponse struct {
    ItemTeamDefinitionPrimaryChannelAllMembersRemovePostResponse
}
// NewItemTeamDefinitionPrimaryChannelAllMembersRemoveResponse instantiates a new ItemTeamDefinitionPrimaryChannelAllMembersRemoveResponse and sets the default values.
func NewItemTeamDefinitionPrimaryChannelAllMembersRemoveResponse()(*ItemTeamDefinitionPrimaryChannelAllMembersRemoveResponse) {
    m := &ItemTeamDefinitionPrimaryChannelAllMembersRemoveResponse{
        ItemTeamDefinitionPrimaryChannelAllMembersRemovePostResponse: *NewItemTeamDefinitionPrimaryChannelAllMembersRemovePostResponse(),
    }
    return m
}
// CreateItemTeamDefinitionPrimaryChannelAllMembersRemoveResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamDefinitionPrimaryChannelAllMembersRemoveResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamDefinitionPrimaryChannelAllMembersRemoveResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemTeamDefinitionPrimaryChannelAllMembersRemovePostResponseable instead.
type ItemTeamDefinitionPrimaryChannelAllMembersRemoveResponseable interface {
    ItemTeamDefinitionPrimaryChannelAllMembersRemovePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
