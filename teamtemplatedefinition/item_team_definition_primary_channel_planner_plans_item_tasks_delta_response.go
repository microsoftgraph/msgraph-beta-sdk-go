// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package teamtemplatedefinition

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemTeamDefinitionPrimaryChannelPlannerPlansItemTasksDeltaGetResponseable instead.
type ItemTeamDefinitionPrimaryChannelPlannerPlansItemTasksDeltaResponse struct {
    ItemTeamDefinitionPrimaryChannelPlannerPlansItemTasksDeltaGetResponse
}
// NewItemTeamDefinitionPrimaryChannelPlannerPlansItemTasksDeltaResponse instantiates a new ItemTeamDefinitionPrimaryChannelPlannerPlansItemTasksDeltaResponse and sets the default values.
func NewItemTeamDefinitionPrimaryChannelPlannerPlansItemTasksDeltaResponse()(*ItemTeamDefinitionPrimaryChannelPlannerPlansItemTasksDeltaResponse) {
    m := &ItemTeamDefinitionPrimaryChannelPlannerPlansItemTasksDeltaResponse{
        ItemTeamDefinitionPrimaryChannelPlannerPlansItemTasksDeltaGetResponse: *NewItemTeamDefinitionPrimaryChannelPlannerPlansItemTasksDeltaGetResponse(),
    }
    return m
}
// CreateItemTeamDefinitionPrimaryChannelPlannerPlansItemTasksDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamDefinitionPrimaryChannelPlannerPlansItemTasksDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamDefinitionPrimaryChannelPlannerPlansItemTasksDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemTeamDefinitionPrimaryChannelPlannerPlansItemTasksDeltaGetResponseable instead.
type ItemTeamDefinitionPrimaryChannelPlannerPlansItemTasksDeltaResponseable interface {
    ItemTeamDefinitionPrimaryChannelPlannerPlansItemTasksDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
