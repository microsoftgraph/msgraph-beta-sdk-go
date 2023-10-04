package teamtemplatedefinition

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemTeamDefinitionChannelsItemMessagesItemRepliesDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemTeamDefinitionChannelsItemMessagesItemRepliesDeltaResponse struct {
    ItemTeamDefinitionChannelsItemMessagesItemRepliesDeltaGetResponse
}
// NewItemTeamDefinitionChannelsItemMessagesItemRepliesDeltaResponse instantiates a new ItemTeamDefinitionChannelsItemMessagesItemRepliesDeltaResponse and sets the default values.
func NewItemTeamDefinitionChannelsItemMessagesItemRepliesDeltaResponse()(*ItemTeamDefinitionChannelsItemMessagesItemRepliesDeltaResponse) {
    m := &ItemTeamDefinitionChannelsItemMessagesItemRepliesDeltaResponse{
        ItemTeamDefinitionChannelsItemMessagesItemRepliesDeltaGetResponse: *NewItemTeamDefinitionChannelsItemMessagesItemRepliesDeltaGetResponse(),
    }
    return m
}
// CreateItemTeamDefinitionChannelsItemMessagesItemRepliesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemTeamDefinitionChannelsItemMessagesItemRepliesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamDefinitionChannelsItemMessagesItemRepliesDeltaResponse(), nil
}
// ItemTeamDefinitionChannelsItemMessagesItemRepliesDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemTeamDefinitionChannelsItemMessagesItemRepliesDeltaResponseable interface {
    ItemTeamDefinitionChannelsItemMessagesItemRepliesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
