// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package teamtemplatedefinition

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemTeamDefinitionChannelsItemAllMembersAddPostResponseable instead.
type ItemTeamDefinitionChannelsItemAllMembersAddResponse struct {
    ItemTeamDefinitionChannelsItemAllMembersAddPostResponse
}
// NewItemTeamDefinitionChannelsItemAllMembersAddResponse instantiates a new ItemTeamDefinitionChannelsItemAllMembersAddResponse and sets the default values.
func NewItemTeamDefinitionChannelsItemAllMembersAddResponse()(*ItemTeamDefinitionChannelsItemAllMembersAddResponse) {
    m := &ItemTeamDefinitionChannelsItemAllMembersAddResponse{
        ItemTeamDefinitionChannelsItemAllMembersAddPostResponse: *NewItemTeamDefinitionChannelsItemAllMembersAddPostResponse(),
    }
    return m
}
// CreateItemTeamDefinitionChannelsItemAllMembersAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamDefinitionChannelsItemAllMembersAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamDefinitionChannelsItemAllMembersAddResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemTeamDefinitionChannelsItemAllMembersAddPostResponseable instead.
type ItemTeamDefinitionChannelsItemAllMembersAddResponseable interface {
    ItemTeamDefinitionChannelsItemAllMembersAddPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
