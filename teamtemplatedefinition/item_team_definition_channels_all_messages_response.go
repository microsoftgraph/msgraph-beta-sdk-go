package teamtemplatedefinition

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemTeamDefinitionChannelsAllMessagesResponse 
// Deprecated: This class is obsolete. Use allMessagesGetResponse instead.
type ItemTeamDefinitionChannelsAllMessagesResponse struct {
    ItemTeamDefinitionChannelsAllMessagesGetResponse
}
// NewItemTeamDefinitionChannelsAllMessagesResponse instantiates a new ItemTeamDefinitionChannelsAllMessagesResponse and sets the default values.
func NewItemTeamDefinitionChannelsAllMessagesResponse()(*ItemTeamDefinitionChannelsAllMessagesResponse) {
    m := &ItemTeamDefinitionChannelsAllMessagesResponse{
        ItemTeamDefinitionChannelsAllMessagesGetResponse: *NewItemTeamDefinitionChannelsAllMessagesGetResponse(),
    }
    return m
}
// CreateItemTeamDefinitionChannelsAllMessagesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemTeamDefinitionChannelsAllMessagesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamDefinitionChannelsAllMessagesResponse(), nil
}
// ItemTeamDefinitionChannelsAllMessagesResponseable 
// Deprecated: This class is obsolete. Use allMessagesGetResponse instead.
type ItemTeamDefinitionChannelsAllMessagesResponseable interface {
    ItemTeamDefinitionChannelsAllMessagesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
