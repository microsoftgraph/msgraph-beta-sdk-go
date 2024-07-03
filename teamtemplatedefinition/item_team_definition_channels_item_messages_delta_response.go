package teamtemplatedefinition

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemTeamDefinitionChannelsItemMessagesDeltaGetResponseable instead.
type ItemTeamDefinitionChannelsItemMessagesDeltaResponse struct {
    ItemTeamDefinitionChannelsItemMessagesDeltaGetResponse
}
// NewItemTeamDefinitionChannelsItemMessagesDeltaResponse instantiates a new ItemTeamDefinitionChannelsItemMessagesDeltaResponse and sets the default values.
func NewItemTeamDefinitionChannelsItemMessagesDeltaResponse()(*ItemTeamDefinitionChannelsItemMessagesDeltaResponse) {
    m := &ItemTeamDefinitionChannelsItemMessagesDeltaResponse{
        ItemTeamDefinitionChannelsItemMessagesDeltaGetResponse: *NewItemTeamDefinitionChannelsItemMessagesDeltaGetResponse(),
    }
    return m
}
// CreateItemTeamDefinitionChannelsItemMessagesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamDefinitionChannelsItemMessagesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamDefinitionChannelsItemMessagesDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemTeamDefinitionChannelsItemMessagesDeltaGetResponseable instead.
type ItemTeamDefinitionChannelsItemMessagesDeltaResponseable interface {
    ItemTeamDefinitionChannelsItemMessagesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
