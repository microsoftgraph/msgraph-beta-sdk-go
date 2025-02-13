package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemTeamPrimaryChannelPlannerPlansDeltaGetResponseable instead.
type ItemTeamPrimaryChannelPlannerPlansDeltaResponse struct {
    ItemTeamPrimaryChannelPlannerPlansDeltaGetResponse
}
// NewItemTeamPrimaryChannelPlannerPlansDeltaResponse instantiates a new ItemTeamPrimaryChannelPlannerPlansDeltaResponse and sets the default values.
func NewItemTeamPrimaryChannelPlannerPlansDeltaResponse()(*ItemTeamPrimaryChannelPlannerPlansDeltaResponse) {
    m := &ItemTeamPrimaryChannelPlannerPlansDeltaResponse{
        ItemTeamPrimaryChannelPlannerPlansDeltaGetResponse: *NewItemTeamPrimaryChannelPlannerPlansDeltaGetResponse(),
    }
    return m
}
// CreateItemTeamPrimaryChannelPlannerPlansDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamPrimaryChannelPlannerPlansDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamPrimaryChannelPlannerPlansDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemTeamPrimaryChannelPlannerPlansDeltaGetResponseable instead.
type ItemTeamPrimaryChannelPlannerPlansDeltaResponseable interface {
    ItemTeamPrimaryChannelPlannerPlansDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
