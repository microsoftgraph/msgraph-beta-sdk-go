package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemChannelsItemPlannerPlansDeltaGetResponseable instead.
type ItemChannelsItemPlannerPlansDeltaResponse struct {
    ItemChannelsItemPlannerPlansDeltaGetResponse
}
// NewItemChannelsItemPlannerPlansDeltaResponse instantiates a new ItemChannelsItemPlannerPlansDeltaResponse and sets the default values.
func NewItemChannelsItemPlannerPlansDeltaResponse()(*ItemChannelsItemPlannerPlansDeltaResponse) {
    m := &ItemChannelsItemPlannerPlansDeltaResponse{
        ItemChannelsItemPlannerPlansDeltaGetResponse: *NewItemChannelsItemPlannerPlansDeltaGetResponse(),
    }
    return m
}
// CreateItemChannelsItemPlannerPlansDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemChannelsItemPlannerPlansDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemChannelsItemPlannerPlansDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemChannelsItemPlannerPlansDeltaGetResponseable instead.
type ItemChannelsItemPlannerPlansDeltaResponseable interface {
    ItemChannelsItemPlannerPlansDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
