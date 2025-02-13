package teamwork

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use DeletedTeamsItemChannelsItemPlannerPlansItemTasksDeltaGetResponseable instead.
type DeletedTeamsItemChannelsItemPlannerPlansItemTasksDeltaResponse struct {
    DeletedTeamsItemChannelsItemPlannerPlansItemTasksDeltaGetResponse
}
// NewDeletedTeamsItemChannelsItemPlannerPlansItemTasksDeltaResponse instantiates a new DeletedTeamsItemChannelsItemPlannerPlansItemTasksDeltaResponse and sets the default values.
func NewDeletedTeamsItemChannelsItemPlannerPlansItemTasksDeltaResponse()(*DeletedTeamsItemChannelsItemPlannerPlansItemTasksDeltaResponse) {
    m := &DeletedTeamsItemChannelsItemPlannerPlansItemTasksDeltaResponse{
        DeletedTeamsItemChannelsItemPlannerPlansItemTasksDeltaGetResponse: *NewDeletedTeamsItemChannelsItemPlannerPlansItemTasksDeltaGetResponse(),
    }
    return m
}
// CreateDeletedTeamsItemChannelsItemPlannerPlansItemTasksDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeletedTeamsItemChannelsItemPlannerPlansItemTasksDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeletedTeamsItemChannelsItemPlannerPlansItemTasksDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use DeletedTeamsItemChannelsItemPlannerPlansItemTasksDeltaGetResponseable instead.
type DeletedTeamsItemChannelsItemPlannerPlansItemTasksDeltaResponseable interface {
    DeletedTeamsItemChannelsItemPlannerPlansItemTasksDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
