package teamtemplatedefinition

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemTeamDefinitionPrimaryChannelGetAllMembersRemovePostResponseable instead.
type ItemTeamDefinitionPrimaryChannelGetAllMembersRemoveResponse struct {
    ItemTeamDefinitionPrimaryChannelGetAllMembersRemovePostResponse
}
// NewItemTeamDefinitionPrimaryChannelGetAllMembersRemoveResponse instantiates a new ItemTeamDefinitionPrimaryChannelGetAllMembersRemoveResponse and sets the default values.
func NewItemTeamDefinitionPrimaryChannelGetAllMembersRemoveResponse()(*ItemTeamDefinitionPrimaryChannelGetAllMembersRemoveResponse) {
    m := &ItemTeamDefinitionPrimaryChannelGetAllMembersRemoveResponse{
        ItemTeamDefinitionPrimaryChannelGetAllMembersRemovePostResponse: *NewItemTeamDefinitionPrimaryChannelGetAllMembersRemovePostResponse(),
    }
    return m
}
// CreateItemTeamDefinitionPrimaryChannelGetAllMembersRemoveResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamDefinitionPrimaryChannelGetAllMembersRemoveResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamDefinitionPrimaryChannelGetAllMembersRemoveResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemTeamDefinitionPrimaryChannelGetAllMembersRemovePostResponseable instead.
type ItemTeamDefinitionPrimaryChannelGetAllMembersRemoveResponseable interface {
    ItemTeamDefinitionPrimaryChannelGetAllMembersRemovePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
