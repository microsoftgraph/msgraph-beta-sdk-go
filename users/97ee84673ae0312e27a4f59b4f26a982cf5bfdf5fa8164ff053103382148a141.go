// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnGetResponseable instead.
type ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnResponse struct {
    ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnGetResponse
}
// NewItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnResponse instantiates a new ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnResponse and sets the default values.
func NewItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnResponse()(*ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnResponse) {
    m := &ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnResponse{
        ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnGetResponse: *NewItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnGetResponse(),
    }
    return m
}
// CreateItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnGetResponseable instead.
type ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnResponseable interface {
    ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
