package planner

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlansItemBucketsItemTasksDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type PlansItemBucketsItemTasksDeltaResponse struct {
    PlansItemBucketsItemTasksDeltaGetResponse
}
// NewPlansItemBucketsItemTasksDeltaResponse instantiates a new PlansItemBucketsItemTasksDeltaResponse and sets the default values.
func NewPlansItemBucketsItemTasksDeltaResponse()(*PlansItemBucketsItemTasksDeltaResponse) {
    m := &PlansItemBucketsItemTasksDeltaResponse{
        PlansItemBucketsItemTasksDeltaGetResponse: *NewPlansItemBucketsItemTasksDeltaGetResponse(),
    }
    return m
}
// CreatePlansItemBucketsItemTasksDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlansItemBucketsItemTasksDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlansItemBucketsItemTasksDeltaResponse(), nil
}
// PlansItemBucketsItemTasksDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type PlansItemBucketsItemTasksDeltaResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PlansItemBucketsItemTasksDeltaGetResponseable
}
