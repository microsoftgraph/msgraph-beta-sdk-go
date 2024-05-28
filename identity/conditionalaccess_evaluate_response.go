package identity

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ConditionalaccessEvaluatePostResponseable instead.
type ConditionalaccessEvaluateResponse struct {
    ConditionalaccessEvaluatePostResponse
}
// NewConditionalaccessEvaluateResponse instantiates a new ConditionalaccessEvaluateResponse and sets the default values.
func NewConditionalaccessEvaluateResponse()(*ConditionalaccessEvaluateResponse) {
    m := &ConditionalaccessEvaluateResponse{
        ConditionalaccessEvaluatePostResponse: *NewConditionalaccessEvaluatePostResponse(),
    }
    return m
}
// CreateConditionalaccessEvaluateResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateConditionalaccessEvaluateResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConditionalaccessEvaluateResponse(), nil
}
// Deprecated: This class is obsolete. Use ConditionalaccessEvaluatePostResponseable instead.
type ConditionalaccessEvaluateResponseable interface {
    ConditionalaccessEvaluatePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
