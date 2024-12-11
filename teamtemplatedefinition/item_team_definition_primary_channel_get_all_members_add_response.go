package teamtemplatedefinition

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemTeamDefinitionPrimaryChannelGetAllMembersAddPostResponseable instead.
type ItemTeamDefinitionPrimaryChannelGetAllMembersAddResponse struct {
    ItemTeamDefinitionPrimaryChannelGetAllMembersAddPostResponse
}
// NewItemTeamDefinitionPrimaryChannelGetAllMembersAddResponse instantiates a new ItemTeamDefinitionPrimaryChannelGetAllMembersAddResponse and sets the default values.
func NewItemTeamDefinitionPrimaryChannelGetAllMembersAddResponse()(*ItemTeamDefinitionPrimaryChannelGetAllMembersAddResponse) {
    m := &ItemTeamDefinitionPrimaryChannelGetAllMembersAddResponse{
        ItemTeamDefinitionPrimaryChannelGetAllMembersAddPostResponse: *NewItemTeamDefinitionPrimaryChannelGetAllMembersAddPostResponse(),
    }
    return m
}
// CreateItemTeamDefinitionPrimaryChannelGetAllMembersAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamDefinitionPrimaryChannelGetAllMembersAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamDefinitionPrimaryChannelGetAllMembersAddResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemTeamDefinitionPrimaryChannelGetAllMembersAddPostResponseable instead.
type ItemTeamDefinitionPrimaryChannelGetAllMembersAddResponseable interface {
    ItemTeamDefinitionPrimaryChannelGetAllMembersAddPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
