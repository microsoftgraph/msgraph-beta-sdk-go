package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemPlannerPlansItemBucketsDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemPlannerPlansItemBucketsDeltaResponse struct {
    ItemPlannerPlansItemBucketsDeltaGetResponse
}
// NewItemPlannerPlansItemBucketsDeltaResponse instantiates a new ItemPlannerPlansItemBucketsDeltaResponse and sets the default values.
func NewItemPlannerPlansItemBucketsDeltaResponse()(*ItemPlannerPlansItemBucketsDeltaResponse) {
    m := &ItemPlannerPlansItemBucketsDeltaResponse{
        ItemPlannerPlansItemBucketsDeltaGetResponse: *NewItemPlannerPlansItemBucketsDeltaGetResponse(),
    }
    return m
}
// CreateItemPlannerPlansItemBucketsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemPlannerPlansItemBucketsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPlannerPlansItemBucketsDeltaResponse(), nil
}
// ItemPlannerPlansItemBucketsDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemPlannerPlansItemBucketsDeltaResponseable interface {
    ItemPlannerPlansItemBucketsDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
