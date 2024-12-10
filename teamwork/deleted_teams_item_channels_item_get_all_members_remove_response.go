package teamwork

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use DeletedTeamsItemChannelsItemGetAllMembersRemovePostResponseable instead.
type DeletedTeamsItemChannelsItemGetAllMembersRemoveResponse struct {
    DeletedTeamsItemChannelsItemGetAllMembersRemovePostResponse
}
// NewDeletedTeamsItemChannelsItemGetAllMembersRemoveResponse instantiates a new DeletedTeamsItemChannelsItemGetAllMembersRemoveResponse and sets the default values.
func NewDeletedTeamsItemChannelsItemGetAllMembersRemoveResponse()(*DeletedTeamsItemChannelsItemGetAllMembersRemoveResponse) {
    m := &DeletedTeamsItemChannelsItemGetAllMembersRemoveResponse{
        DeletedTeamsItemChannelsItemGetAllMembersRemovePostResponse: *NewDeletedTeamsItemChannelsItemGetAllMembersRemovePostResponse(),
    }
    return m
}
// CreateDeletedTeamsItemChannelsItemGetAllMembersRemoveResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeletedTeamsItemChannelsItemGetAllMembersRemoveResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeletedTeamsItemChannelsItemGetAllMembersRemoveResponse(), nil
}
// Deprecated: This class is obsolete. Use DeletedTeamsItemChannelsItemGetAllMembersRemovePostResponseable instead.
type DeletedTeamsItemChannelsItemGetAllMembersRemoveResponseable interface {
    DeletedTeamsItemChannelsItemGetAllMembersRemovePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
