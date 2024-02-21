package planner

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use PlansItemBucketsDeltaGetResponseable instead.
type PlansItemBucketsDeltaResponse struct {
    PlansItemBucketsDeltaGetResponse
}
// NewPlansItemBucketsDeltaResponse instantiates a new PlansItemBucketsDeltaResponse and sets the default values.
func NewPlansItemBucketsDeltaResponse()(*PlansItemBucketsDeltaResponse) {
    m := &PlansItemBucketsDeltaResponse{
        PlansItemBucketsDeltaGetResponse: *NewPlansItemBucketsDeltaGetResponse(),
    }
    return m
}
// CreatePlansItemBucketsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePlansItemBucketsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlansItemBucketsDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use PlansItemBucketsDeltaGetResponseable instead.
type PlansItemBucketsDeltaResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PlansItemBucketsDeltaGetResponseable
}
