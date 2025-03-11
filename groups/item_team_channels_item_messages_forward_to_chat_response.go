package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemTeamChannelsItemMessagesForwardToChatPostResponseable instead.
type ItemTeamChannelsItemMessagesForwardToChatResponse struct {
    ItemTeamChannelsItemMessagesForwardToChatPostResponse
}
// NewItemTeamChannelsItemMessagesForwardToChatResponse instantiates a new ItemTeamChannelsItemMessagesForwardToChatResponse and sets the default values.
func NewItemTeamChannelsItemMessagesForwardToChatResponse()(*ItemTeamChannelsItemMessagesForwardToChatResponse) {
    m := &ItemTeamChannelsItemMessagesForwardToChatResponse{
        ItemTeamChannelsItemMessagesForwardToChatPostResponse: *NewItemTeamChannelsItemMessagesForwardToChatPostResponse(),
    }
    return m
}
// CreateItemTeamChannelsItemMessagesForwardToChatResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamChannelsItemMessagesForwardToChatResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamChannelsItemMessagesForwardToChatResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemTeamChannelsItemMessagesForwardToChatPostResponseable instead.
type ItemTeamChannelsItemMessagesForwardToChatResponseable interface {
    ItemTeamChannelsItemMessagesForwardToChatPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
