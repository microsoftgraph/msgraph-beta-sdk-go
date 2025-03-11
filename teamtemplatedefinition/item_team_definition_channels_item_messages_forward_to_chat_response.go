package teamtemplatedefinition

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemTeamDefinitionChannelsItemMessagesForwardToChatPostResponseable instead.
type ItemTeamDefinitionChannelsItemMessagesForwardToChatResponse struct {
    ItemTeamDefinitionChannelsItemMessagesForwardToChatPostResponse
}
// NewItemTeamDefinitionChannelsItemMessagesForwardToChatResponse instantiates a new ItemTeamDefinitionChannelsItemMessagesForwardToChatResponse and sets the default values.
func NewItemTeamDefinitionChannelsItemMessagesForwardToChatResponse()(*ItemTeamDefinitionChannelsItemMessagesForwardToChatResponse) {
    m := &ItemTeamDefinitionChannelsItemMessagesForwardToChatResponse{
        ItemTeamDefinitionChannelsItemMessagesForwardToChatPostResponse: *NewItemTeamDefinitionChannelsItemMessagesForwardToChatPostResponse(),
    }
    return m
}
// CreateItemTeamDefinitionChannelsItemMessagesForwardToChatResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamDefinitionChannelsItemMessagesForwardToChatResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamDefinitionChannelsItemMessagesForwardToChatResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemTeamDefinitionChannelsItemMessagesForwardToChatPostResponseable instead.
type ItemTeamDefinitionChannelsItemMessagesForwardToChatResponseable interface {
    ItemTeamDefinitionChannelsItemMessagesForwardToChatPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
