package teamtemplatedefinition

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemTeamDefinitionMembersAddResponse 
// Deprecated: This class is obsolete. Use addPostResponse instead.
type ItemTeamDefinitionMembersAddResponse struct {
    ItemTeamDefinitionMembersAddPostResponse
}
// NewItemTeamDefinitionMembersAddResponse instantiates a new ItemTeamDefinitionMembersAddResponse and sets the default values.
func NewItemTeamDefinitionMembersAddResponse()(*ItemTeamDefinitionMembersAddResponse) {
    m := &ItemTeamDefinitionMembersAddResponse{
        ItemTeamDefinitionMembersAddPostResponse: *NewItemTeamDefinitionMembersAddPostResponse(),
    }
    return m
}
// CreateItemTeamDefinitionMembersAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemTeamDefinitionMembersAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamDefinitionMembersAddResponse(), nil
}
// ItemTeamDefinitionMembersAddResponseable 
// Deprecated: This class is obsolete. Use addPostResponse instead.
type ItemTeamDefinitionMembersAddResponseable interface {
    ItemTeamDefinitionMembersAddPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
