// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemChannelsAllMessagesGetResponseable instead.
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
// returns a Parsable when successful
func CreateItemChannelsAllMessagesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemChannelsAllMessagesResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemChannelsAllMessagesGetResponseable instead.
type ItemChannelsAllMessagesResponseable interface {
    ItemChannelsAllMessagesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
