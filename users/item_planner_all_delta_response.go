package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemPlannerAllDeltaGetResponseable instead.
type ItemPlannerAllDeltaResponse struct {
    ItemPlannerAllDeltaGetResponse
}
// NewItemPlannerAllDeltaResponse instantiates a new ItemPlannerAllDeltaResponse and sets the default values.
func NewItemPlannerAllDeltaResponse()(*ItemPlannerAllDeltaResponse) {
    m := &ItemPlannerAllDeltaResponse{
        ItemPlannerAllDeltaGetResponse: *NewItemPlannerAllDeltaGetResponse(),
    }
    return m
}
// CreateItemPlannerAllDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemPlannerAllDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPlannerAllDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemPlannerAllDeltaGetResponseable instead.
type ItemPlannerAllDeltaResponseable interface {
    ItemPlannerAllDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
