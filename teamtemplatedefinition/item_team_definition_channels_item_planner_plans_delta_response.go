package teamtemplatedefinition

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemTeamDefinitionChannelsItemPlannerPlansDeltaGetResponseable instead.
type ItemTeamDefinitionChannelsItemPlannerPlansDeltaResponse struct {
    ItemTeamDefinitionChannelsItemPlannerPlansDeltaGetResponse
}
// NewItemTeamDefinitionChannelsItemPlannerPlansDeltaResponse instantiates a new ItemTeamDefinitionChannelsItemPlannerPlansDeltaResponse and sets the default values.
func NewItemTeamDefinitionChannelsItemPlannerPlansDeltaResponse()(*ItemTeamDefinitionChannelsItemPlannerPlansDeltaResponse) {
    m := &ItemTeamDefinitionChannelsItemPlannerPlansDeltaResponse{
        ItemTeamDefinitionChannelsItemPlannerPlansDeltaGetResponse: *NewItemTeamDefinitionChannelsItemPlannerPlansDeltaGetResponse(),
    }
    return m
}
// CreateItemTeamDefinitionChannelsItemPlannerPlansDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamDefinitionChannelsItemPlannerPlansDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamDefinitionChannelsItemPlannerPlansDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemTeamDefinitionChannelsItemPlannerPlansDeltaGetResponseable instead.
type ItemTeamDefinitionChannelsItemPlannerPlansDeltaResponseable interface {
    ItemTeamDefinitionChannelsItemPlannerPlansDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
