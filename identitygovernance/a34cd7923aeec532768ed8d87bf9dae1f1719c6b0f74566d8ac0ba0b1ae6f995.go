// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsFilterByCurrentUserWithOnGetResponseable instead.
type AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsFilterByCurrentUserWithOnResponse struct {
    AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsFilterByCurrentUserWithOnGetResponse
}
// NewAccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsFilterByCurrentUserWithOnResponse instantiates a new AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsFilterByCurrentUserWithOnResponse and sets the default values.
func NewAccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsFilterByCurrentUserWithOnResponse()(*AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsFilterByCurrentUserWithOnResponse) {
    m := &AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsFilterByCurrentUserWithOnResponse{
        AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsFilterByCurrentUserWithOnGetResponse: *NewAccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsFilterByCurrentUserWithOnGetResponse(),
    }
    return m
}
// CreateAccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsFilterByCurrentUserWithOnResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsFilterByCurrentUserWithOnResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsFilterByCurrentUserWithOnResponse(), nil
}
// Deprecated: This class is obsolete. Use AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsFilterByCurrentUserWithOnGetResponseable instead.
type AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsFilterByCurrentUserWithOnResponseable interface {
    AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsFilterByCurrentUserWithOnGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
