package planner

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlansItemTasksDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type PlansItemTasksDeltaResponse struct {
    PlansItemTasksDeltaGetResponse
}
// NewPlansItemTasksDeltaResponse instantiates a new PlansItemTasksDeltaResponse and sets the default values.
func NewPlansItemTasksDeltaResponse()(*PlansItemTasksDeltaResponse) {
    m := &PlansItemTasksDeltaResponse{
        PlansItemTasksDeltaGetResponse: *NewPlansItemTasksDeltaGetResponse(),
    }
    return m
}
// CreatePlansItemTasksDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlansItemTasksDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlansItemTasksDeltaResponse(), nil
}
// PlansItemTasksDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type PlansItemTasksDeltaResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PlansItemTasksDeltaGetResponseable
}
