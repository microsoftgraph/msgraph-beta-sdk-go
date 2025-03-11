package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemChannelsItemMessagesItemRepliesForwardToChatPostResponseable instead.
type ItemChannelsItemMessagesItemRepliesForwardToChatResponse struct {
    ItemChannelsItemMessagesItemRepliesForwardToChatPostResponse
}
// NewItemChannelsItemMessagesItemRepliesForwardToChatResponse instantiates a new ItemChannelsItemMessagesItemRepliesForwardToChatResponse and sets the default values.
func NewItemChannelsItemMessagesItemRepliesForwardToChatResponse()(*ItemChannelsItemMessagesItemRepliesForwardToChatResponse) {
    m := &ItemChannelsItemMessagesItemRepliesForwardToChatResponse{
        ItemChannelsItemMessagesItemRepliesForwardToChatPostResponse: *NewItemChannelsItemMessagesItemRepliesForwardToChatPostResponse(),
    }
    return m
}
// CreateItemChannelsItemMessagesItemRepliesForwardToChatResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemChannelsItemMessagesItemRepliesForwardToChatResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemChannelsItemMessagesItemRepliesForwardToChatResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemChannelsItemMessagesItemRepliesForwardToChatPostResponseable instead.
type ItemChannelsItemMessagesItemRepliesForwardToChatResponseable interface {
    ItemChannelsItemMessagesItemRepliesForwardToChatPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
