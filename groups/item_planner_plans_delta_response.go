package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemPlannerPlansDeltaGetResponseable instead.
type ItemPlannerPlansDeltaResponse struct {
    ItemPlannerPlansDeltaGetResponse
}
// NewItemPlannerPlansDeltaResponse instantiates a new ItemPlannerPlansDeltaResponse and sets the default values.
func NewItemPlannerPlansDeltaResponse()(*ItemPlannerPlansDeltaResponse) {
    m := &ItemPlannerPlansDeltaResponse{
        ItemPlannerPlansDeltaGetResponse: *NewItemPlannerPlansDeltaGetResponse(),
    }
    return m
}
// CreateItemPlannerPlansDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemPlannerPlansDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPlannerPlansDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemPlannerPlansDeltaGetResponseable instead.
type ItemPlannerPlansDeltaResponseable interface {
    ItemPlannerPlansDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
