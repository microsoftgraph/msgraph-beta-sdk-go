package teamtemplatedefinition

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemTeamDefinitionMembersRemovePostResponseable instead.
type ItemTeamDefinitionMembersRemoveResponse struct {
    ItemTeamDefinitionMembersRemovePostResponse
}
// NewItemTeamDefinitionMembersRemoveResponse instantiates a new ItemTeamDefinitionMembersRemoveResponse and sets the default values.
func NewItemTeamDefinitionMembersRemoveResponse()(*ItemTeamDefinitionMembersRemoveResponse) {
    m := &ItemTeamDefinitionMembersRemoveResponse{
        ItemTeamDefinitionMembersRemovePostResponse: *NewItemTeamDefinitionMembersRemovePostResponse(),
    }
    return m
}
// CreateItemTeamDefinitionMembersRemoveResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamDefinitionMembersRemoveResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamDefinitionMembersRemoveResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemTeamDefinitionMembersRemovePostResponseable instead.
type ItemTeamDefinitionMembersRemoveResponseable interface {
    ItemTeamDefinitionMembersRemovePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
