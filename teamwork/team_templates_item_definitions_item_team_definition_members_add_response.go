package teamwork

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamTemplatesItemDefinitionsItemTeamDefinitionMembersAddResponse 
// Deprecated: This class is obsolete. Use addPostResponse instead.
type TeamTemplatesItemDefinitionsItemTeamDefinitionMembersAddResponse struct {
    TeamTemplatesItemDefinitionsItemTeamDefinitionMembersAddPostResponse
}
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionMembersAddResponse instantiates a new TeamTemplatesItemDefinitionsItemTeamDefinitionMembersAddResponse and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionMembersAddResponse()(*TeamTemplatesItemDefinitionsItemTeamDefinitionMembersAddResponse) {
    m := &TeamTemplatesItemDefinitionsItemTeamDefinitionMembersAddResponse{
        TeamTemplatesItemDefinitionsItemTeamDefinitionMembersAddPostResponse: *NewTeamTemplatesItemDefinitionsItemTeamDefinitionMembersAddPostResponse(),
    }
    return m
}
// CreateTeamTemplatesItemDefinitionsItemTeamDefinitionMembersAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamTemplatesItemDefinitionsItemTeamDefinitionMembersAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionMembersAddResponse(), nil
}
// TeamTemplatesItemDefinitionsItemTeamDefinitionMembersAddResponseable 
// Deprecated: This class is obsolete. Use addPostResponse instead.
type TeamTemplatesItemDefinitionsItemTeamDefinitionMembersAddResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TeamTemplatesItemDefinitionsItemTeamDefinitionMembersAddPostResponseable
}
