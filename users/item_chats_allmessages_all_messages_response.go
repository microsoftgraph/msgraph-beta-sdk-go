package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemChatsAllmessagesAllMessagesGetResponseable instead.
type ItemChatsAllmessagesAllMessagesResponse struct {
    ItemChatsAllmessagesAllMessagesGetResponse
}
// NewItemChatsAllmessagesAllMessagesResponse instantiates a new ItemChatsAllmessagesAllMessagesResponse and sets the default values.
func NewItemChatsAllmessagesAllMessagesResponse()(*ItemChatsAllmessagesAllMessagesResponse) {
    m := &ItemChatsAllmessagesAllMessagesResponse{
        ItemChatsAllmessagesAllMessagesGetResponse: *NewItemChatsAllmessagesAllMessagesGetResponse(),
    }
    return m
}
// CreateItemChatsAllmessagesAllMessagesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemChatsAllmessagesAllMessagesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemChatsAllmessagesAllMessagesResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemChatsAllmessagesAllMessagesGetResponseable instead.
type ItemChatsAllmessagesAllMessagesResponseable interface {
    ItemChatsAllmessagesAllMessagesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
