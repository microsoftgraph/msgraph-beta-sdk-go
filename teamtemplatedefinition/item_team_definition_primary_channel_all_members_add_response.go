package teamtemplatedefinition

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemTeamDefinitionPrimaryChannelAllMembersAddPostResponseable instead.
type ItemTeamDefinitionPrimaryChannelAllMembersAddResponse struct {
    ItemTeamDefinitionPrimaryChannelAllMembersAddPostResponse
}
// NewItemTeamDefinitionPrimaryChannelAllMembersAddResponse instantiates a new ItemTeamDefinitionPrimaryChannelAllMembersAddResponse and sets the default values.
func NewItemTeamDefinitionPrimaryChannelAllMembersAddResponse()(*ItemTeamDefinitionPrimaryChannelAllMembersAddResponse) {
    m := &ItemTeamDefinitionPrimaryChannelAllMembersAddResponse{
        ItemTeamDefinitionPrimaryChannelAllMembersAddPostResponse: *NewItemTeamDefinitionPrimaryChannelAllMembersAddPostResponse(),
    }
    return m
}
// CreateItemTeamDefinitionPrimaryChannelAllMembersAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamDefinitionPrimaryChannelAllMembersAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamDefinitionPrimaryChannelAllMembersAddResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemTeamDefinitionPrimaryChannelAllMembersAddPostResponseable instead.
type ItemTeamDefinitionPrimaryChannelAllMembersAddResponseable interface {
    ItemTeamDefinitionPrimaryChannelAllMembersAddPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
