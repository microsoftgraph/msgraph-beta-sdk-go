package teamtemplatedefinition

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type ItemTeamDefinitionChannelsGetAllRetainedMessagesResponse struct {
    ItemTeamDefinitionChannelsGetAllRetainedMessagesGetResponse
}
// NewItemTeamDefinitionChannelsGetAllRetainedMessagesResponse instantiates a new ItemTeamDefinitionChannelsGetAllRetainedMessagesResponse and sets the default values.
func NewItemTeamDefinitionChannelsGetAllRetainedMessagesResponse()(*ItemTeamDefinitionChannelsGetAllRetainedMessagesResponse) {
    m := &ItemTeamDefinitionChannelsGetAllRetainedMessagesResponse{
        ItemTeamDefinitionChannelsGetAllRetainedMessagesGetResponse: *NewItemTeamDefinitionChannelsGetAllRetainedMessagesGetResponse(),
    }
    return m
}
// CreateItemTeamDefinitionChannelsGetAllRetainedMessagesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamDefinitionChannelsGetAllRetainedMessagesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamDefinitionChannelsGetAllRetainedMessagesResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type ItemTeamDefinitionChannelsGetAllRetainedMessagesResponseable interface {
    ItemTeamDefinitionChannelsGetAllRetainedMessagesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
