package teamtemplatedefinition

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemTeamdefinitionChannelsAllmessagesAllMessagesGetResponseable instead.
type ItemTeamdefinitionChannelsAllmessagesAllMessagesResponse struct {
    ItemTeamdefinitionChannelsAllmessagesAllMessagesGetResponse
}
// NewItemTeamdefinitionChannelsAllmessagesAllMessagesResponse instantiates a new ItemTeamdefinitionChannelsAllmessagesAllMessagesResponse and sets the default values.
func NewItemTeamdefinitionChannelsAllmessagesAllMessagesResponse()(*ItemTeamdefinitionChannelsAllmessagesAllMessagesResponse) {
    m := &ItemTeamdefinitionChannelsAllmessagesAllMessagesResponse{
        ItemTeamdefinitionChannelsAllmessagesAllMessagesGetResponse: *NewItemTeamdefinitionChannelsAllmessagesAllMessagesGetResponse(),
    }
    return m
}
// CreateItemTeamdefinitionChannelsAllmessagesAllMessagesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamdefinitionChannelsAllmessagesAllMessagesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamdefinitionChannelsAllmessagesAllMessagesResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemTeamdefinitionChannelsAllmessagesAllMessagesGetResponseable instead.
type ItemTeamdefinitionChannelsAllmessagesAllMessagesResponseable interface {
    ItemTeamdefinitionChannelsAllmessagesAllMessagesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
