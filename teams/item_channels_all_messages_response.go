package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemChannelsAllMessagesResponse 
// Deprecated: This class is obsolete. Use allMessagesGetResponse instead.
type ItemChannelsAllMessagesResponse struct {
    ItemChannelsAllMessagesGetResponse
}
// NewItemChannelsAllMessagesResponse instantiates a new ItemChannelsAllMessagesResponse and sets the default values.
func NewItemChannelsAllMessagesResponse()(*ItemChannelsAllMessagesResponse) {
    m := &ItemChannelsAllMessagesResponse{
        ItemChannelsAllMessagesGetResponse: *NewItemChannelsAllMessagesGetResponse(),
    }
    return m
}
// CreateItemChannelsAllMessagesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemChannelsAllMessagesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemChannelsAllMessagesResponse(), nil
}
// ItemChannelsAllMessagesResponseable 
// Deprecated: This class is obsolete. Use allMessagesGetResponse instead.
type ItemChannelsAllMessagesResponseable interface {
    ItemChannelsAllMessagesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
