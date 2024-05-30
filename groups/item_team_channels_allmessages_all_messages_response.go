package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemTeamChannelsAllmessagesAllMessagesGetResponseable instead.
type ItemTeamChannelsAllmessagesAllMessagesResponse struct {
    ItemTeamChannelsAllmessagesAllMessagesGetResponse
}
// NewItemTeamChannelsAllmessagesAllMessagesResponse instantiates a new ItemTeamChannelsAllmessagesAllMessagesResponse and sets the default values.
func NewItemTeamChannelsAllmessagesAllMessagesResponse()(*ItemTeamChannelsAllmessagesAllMessagesResponse) {
    m := &ItemTeamChannelsAllmessagesAllMessagesResponse{
        ItemTeamChannelsAllmessagesAllMessagesGetResponse: *NewItemTeamChannelsAllmessagesAllMessagesGetResponse(),
    }
    return m
}
// CreateItemTeamChannelsAllmessagesAllMessagesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamChannelsAllmessagesAllMessagesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamChannelsAllmessagesAllMessagesResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemTeamChannelsAllmessagesAllMessagesGetResponseable instead.
type ItemTeamChannelsAllmessagesAllMessagesResponseable interface {
    ItemTeamChannelsAllmessagesAllMessagesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
