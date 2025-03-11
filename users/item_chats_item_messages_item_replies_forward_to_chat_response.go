package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemChatsItemMessagesItemRepliesForwardToChatPostResponseable instead.
type ItemChatsItemMessagesItemRepliesForwardToChatResponse struct {
    ItemChatsItemMessagesItemRepliesForwardToChatPostResponse
}
// NewItemChatsItemMessagesItemRepliesForwardToChatResponse instantiates a new ItemChatsItemMessagesItemRepliesForwardToChatResponse and sets the default values.
func NewItemChatsItemMessagesItemRepliesForwardToChatResponse()(*ItemChatsItemMessagesItemRepliesForwardToChatResponse) {
    m := &ItemChatsItemMessagesItemRepliesForwardToChatResponse{
        ItemChatsItemMessagesItemRepliesForwardToChatPostResponse: *NewItemChatsItemMessagesItemRepliesForwardToChatPostResponse(),
    }
    return m
}
// CreateItemChatsItemMessagesItemRepliesForwardToChatResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemChatsItemMessagesItemRepliesForwardToChatResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemChatsItemMessagesItemRepliesForwardToChatResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemChatsItemMessagesItemRepliesForwardToChatPostResponseable instead.
type ItemChatsItemMessagesItemRepliesForwardToChatResponseable interface {
    ItemChatsItemMessagesItemRepliesForwardToChatPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
