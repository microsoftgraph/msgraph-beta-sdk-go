package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type AccessReviewsDecisionsItemInstanceStagesFilterByCurrentUserWithOnResponse struct {
    AccessReviewsDecisionsItemInstanceStagesFilterByCurrentUserWithOnGetResponse
}
// NewAccessReviewsDecisionsItemInstanceStagesFilterByCurrentUserWithOnResponse instantiates a new AccessReviewsDecisionsItemInstanceStagesFilterByCurrentUserWithOnResponse and sets the default values.
func NewAccessReviewsDecisionsItemInstanceStagesFilterByCurrentUserWithOnResponse()(*AccessReviewsDecisionsItemInstanceStagesFilterByCurrentUserWithOnResponse) {
    m := &AccessReviewsDecisionsItemInstanceStagesFilterByCurrentUserWithOnResponse{
        AccessReviewsDecisionsItemInstanceStagesFilterByCurrentUserWithOnGetResponse: *NewAccessReviewsDecisionsItemInstanceStagesFilterByCurrentUserWithOnGetResponse(),
    }
    return m
}
// CreateAccessReviewsDecisionsItemInstanceStagesFilterByCurrentUserWithOnResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccessReviewsDecisionsItemInstanceStagesFilterByCurrentUserWithOnResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessReviewsDecisionsItemInstanceStagesFilterByCurrentUserWithOnResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type AccessReviewsDecisionsItemInstanceStagesFilterByCurrentUserWithOnResponseable interface {
    AccessReviewsDecisionsItemInstanceStagesFilterByCurrentUserWithOnGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
