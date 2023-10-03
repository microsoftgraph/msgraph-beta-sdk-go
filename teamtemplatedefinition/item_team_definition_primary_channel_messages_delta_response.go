package teamtemplatedefinition

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemTeamDefinitionPrimaryChannelMessagesDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemTeamDefinitionPrimaryChannelMessagesDeltaResponse struct {
    ItemTeamDefinitionPrimaryChannelMessagesDeltaGetResponse
}
// NewItemTeamDefinitionPrimaryChannelMessagesDeltaResponse instantiates a new ItemTeamDefinitionPrimaryChannelMessagesDeltaResponse and sets the default values.
func NewItemTeamDefinitionPrimaryChannelMessagesDeltaResponse()(*ItemTeamDefinitionPrimaryChannelMessagesDeltaResponse) {
    m := &ItemTeamDefinitionPrimaryChannelMessagesDeltaResponse{
        ItemTeamDefinitionPrimaryChannelMessagesDeltaGetResponse: *NewItemTeamDefinitionPrimaryChannelMessagesDeltaGetResponse(),
    }
    return m
}
// CreateItemTeamDefinitionPrimaryChannelMessagesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemTeamDefinitionPrimaryChannelMessagesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamDefinitionPrimaryChannelMessagesDeltaResponse(), nil
}
// ItemTeamDefinitionPrimaryChannelMessagesDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemTeamDefinitionPrimaryChannelMessagesDeltaResponseable interface {
    ItemTeamDefinitionPrimaryChannelMessagesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
