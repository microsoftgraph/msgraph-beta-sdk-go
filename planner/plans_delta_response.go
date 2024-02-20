package planner

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use PlansDeltaGetResponseable instead.
type PlansDeltaResponse struct {
    PlansDeltaGetResponse
}
// NewPlansDeltaResponse instantiates a new PlansDeltaResponse and sets the default values.
func NewPlansDeltaResponse()(*PlansDeltaResponse) {
    m := &PlansDeltaResponse{
        PlansDeltaGetResponse: *NewPlansDeltaGetResponse(),
    }
    return m
}
// CreatePlansDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePlansDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlansDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use PlansDeltaGetResponseable instead.
type PlansDeltaResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PlansDeltaGetResponseable
}
