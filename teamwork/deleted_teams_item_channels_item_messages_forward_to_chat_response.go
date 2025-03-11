package teamwork

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use DeletedTeamsItemChannelsItemMessagesForwardToChatPostResponseable instead.
type DeletedTeamsItemChannelsItemMessagesForwardToChatResponse struct {
    DeletedTeamsItemChannelsItemMessagesForwardToChatPostResponse
}
// NewDeletedTeamsItemChannelsItemMessagesForwardToChatResponse instantiates a new DeletedTeamsItemChannelsItemMessagesForwardToChatResponse and sets the default values.
func NewDeletedTeamsItemChannelsItemMessagesForwardToChatResponse()(*DeletedTeamsItemChannelsItemMessagesForwardToChatResponse) {
    m := &DeletedTeamsItemChannelsItemMessagesForwardToChatResponse{
        DeletedTeamsItemChannelsItemMessagesForwardToChatPostResponse: *NewDeletedTeamsItemChannelsItemMessagesForwardToChatPostResponse(),
    }
    return m
}
// CreateDeletedTeamsItemChannelsItemMessagesForwardToChatResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeletedTeamsItemChannelsItemMessagesForwardToChatResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeletedTeamsItemChannelsItemMessagesForwardToChatResponse(), nil
}
// Deprecated: This class is obsolete. Use DeletedTeamsItemChannelsItemMessagesForwardToChatPostResponseable instead.
type DeletedTeamsItemChannelsItemMessagesForwardToChatResponseable interface {
    DeletedTeamsItemChannelsItemMessagesForwardToChatPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
