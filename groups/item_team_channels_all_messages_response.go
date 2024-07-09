package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemTeamChannelsAllMessagesGetResponseable instead.
type ItemTeamChannelsAllMessagesResponse struct {
    ItemTeamChannelsAllMessagesGetResponse
}
// NewItemTeamChannelsAllMessagesResponse instantiates a new ItemTeamChannelsAllMessagesResponse and sets the default values.
func NewItemTeamChannelsAllMessagesResponse()(*ItemTeamChannelsAllMessagesResponse) {
    m := &ItemTeamChannelsAllMessagesResponse{
        ItemTeamChannelsAllMessagesGetResponse: *NewItemTeamChannelsAllMessagesGetResponse(),
    }
    return m
}
// CreateItemTeamChannelsAllMessagesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamChannelsAllMessagesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamChannelsAllMessagesResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemTeamChannelsAllMessagesGetResponseable instead.
type ItemTeamChannelsAllMessagesResponseable interface {
    ItemTeamChannelsAllMessagesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
