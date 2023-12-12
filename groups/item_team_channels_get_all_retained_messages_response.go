package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemTeamChannelsGetAllRetainedMessagesResponse 
// Deprecated: This class is obsolete. Use getAllRetainedMessagesGetResponse instead.
type ItemTeamChannelsGetAllRetainedMessagesResponse struct {
    ItemTeamChannelsGetAllRetainedMessagesGetResponse
}
// NewItemTeamChannelsGetAllRetainedMessagesResponse instantiates a new ItemTeamChannelsGetAllRetainedMessagesResponse and sets the default values.
func NewItemTeamChannelsGetAllRetainedMessagesResponse()(*ItemTeamChannelsGetAllRetainedMessagesResponse) {
    m := &ItemTeamChannelsGetAllRetainedMessagesResponse{
        ItemTeamChannelsGetAllRetainedMessagesGetResponse: *NewItemTeamChannelsGetAllRetainedMessagesGetResponse(),
    }
    return m
}
// CreateItemTeamChannelsGetAllRetainedMessagesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemTeamChannelsGetAllRetainedMessagesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamChannelsGetAllRetainedMessagesResponse(), nil
}
// ItemTeamChannelsGetAllRetainedMessagesResponseable 
// Deprecated: This class is obsolete. Use getAllRetainedMessagesGetResponse instead.
type ItemTeamChannelsGetAllRetainedMessagesResponseable interface {
    ItemTeamChannelsGetAllRetainedMessagesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
