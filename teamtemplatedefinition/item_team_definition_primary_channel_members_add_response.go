package teamtemplatedefinition

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type ItemTeamDefinitionPrimaryChannelMembersAddResponse struct {
    ItemTeamDefinitionPrimaryChannelMembersAddPostResponse
}
// NewItemTeamDefinitionPrimaryChannelMembersAddResponse instantiates a new ItemTeamDefinitionPrimaryChannelMembersAddResponse and sets the default values.
func NewItemTeamDefinitionPrimaryChannelMembersAddResponse()(*ItemTeamDefinitionPrimaryChannelMembersAddResponse) {
    m := &ItemTeamDefinitionPrimaryChannelMembersAddResponse{
        ItemTeamDefinitionPrimaryChannelMembersAddPostResponse: *NewItemTeamDefinitionPrimaryChannelMembersAddPostResponse(),
    }
    return m
}
// CreateItemTeamDefinitionPrimaryChannelMembersAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamDefinitionPrimaryChannelMembersAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamDefinitionPrimaryChannelMembersAddResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type ItemTeamDefinitionPrimaryChannelMembersAddResponseable interface {
    ItemTeamDefinitionPrimaryChannelMembersAddPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
