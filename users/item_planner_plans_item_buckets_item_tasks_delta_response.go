package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemPlannerPlansItemBucketsItemTasksDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemPlannerPlansItemBucketsItemTasksDeltaResponse struct {
    ItemPlannerPlansItemBucketsItemTasksDeltaGetResponse
}
// NewItemPlannerPlansItemBucketsItemTasksDeltaResponse instantiates a new ItemPlannerPlansItemBucketsItemTasksDeltaResponse and sets the default values.
func NewItemPlannerPlansItemBucketsItemTasksDeltaResponse()(*ItemPlannerPlansItemBucketsItemTasksDeltaResponse) {
    m := &ItemPlannerPlansItemBucketsItemTasksDeltaResponse{
        ItemPlannerPlansItemBucketsItemTasksDeltaGetResponse: *NewItemPlannerPlansItemBucketsItemTasksDeltaGetResponse(),
    }
    return m
}
// CreateItemPlannerPlansItemBucketsItemTasksDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemPlannerPlansItemBucketsItemTasksDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPlannerPlansItemBucketsItemTasksDeltaResponse(), nil
}
// ItemPlannerPlansItemBucketsItemTasksDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemPlannerPlansItemBucketsItemTasksDeltaResponseable interface {
    ItemPlannerPlansItemBucketsItemTasksDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
