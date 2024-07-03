package teamtemplatedefinition

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemTeamDefinitionChannelsGetAllMessagesGetResponseable instead.
type ItemTeamDefinitionChannelsGetAllMessagesResponse struct {
    ItemTeamDefinitionChannelsGetAllMessagesGetResponse
}
// NewItemTeamDefinitionChannelsGetAllMessagesResponse instantiates a new ItemTeamDefinitionChannelsGetAllMessagesResponse and sets the default values.
func NewItemTeamDefinitionChannelsGetAllMessagesResponse()(*ItemTeamDefinitionChannelsGetAllMessagesResponse) {
    m := &ItemTeamDefinitionChannelsGetAllMessagesResponse{
        ItemTeamDefinitionChannelsGetAllMessagesGetResponse: *NewItemTeamDefinitionChannelsGetAllMessagesGetResponse(),
    }
    return m
}
// CreateItemTeamDefinitionChannelsGetAllMessagesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamDefinitionChannelsGetAllMessagesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamDefinitionChannelsGetAllMessagesResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemTeamDefinitionChannelsGetAllMessagesGetResponseable instead.
type ItemTeamDefinitionChannelsGetAllMessagesResponseable interface {
    ItemTeamDefinitionChannelsGetAllMessagesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
