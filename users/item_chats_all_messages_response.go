// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemChatsAllMessagesGetResponseable instead.
type ItemChatsAllMessagesResponse struct {
    ItemChatsAllMessagesGetResponse
}
// NewItemChatsAllMessagesResponse instantiates a new ItemChatsAllMessagesResponse and sets the default values.
func NewItemChatsAllMessagesResponse()(*ItemChatsAllMessagesResponse) {
    m := &ItemChatsAllMessagesResponse{
        ItemChatsAllMessagesGetResponse: *NewItemChatsAllMessagesGetResponse(),
    }
    return m
}
// CreateItemChatsAllMessagesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemChatsAllMessagesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemChatsAllMessagesResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemChatsAllMessagesGetResponseable instead.
type ItemChatsAllMessagesResponseable interface {
    ItemChatsAllMessagesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
