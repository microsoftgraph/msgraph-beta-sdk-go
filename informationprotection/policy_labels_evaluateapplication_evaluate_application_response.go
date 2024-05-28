package informationprotection

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use PolicyLabelsEvaluateapplicationEvaluateApplicationPostResponseable instead.
type PolicyLabelsEvaluateapplicationEvaluateApplicationResponse struct {
    PolicyLabelsEvaluateapplicationEvaluateApplicationPostResponse
}
// NewPolicyLabelsEvaluateapplicationEvaluateApplicationResponse instantiates a new PolicyLabelsEvaluateapplicationEvaluateApplicationResponse and sets the default values.
func NewPolicyLabelsEvaluateapplicationEvaluateApplicationResponse()(*PolicyLabelsEvaluateapplicationEvaluateApplicationResponse) {
    m := &PolicyLabelsEvaluateapplicationEvaluateApplicationResponse{
        PolicyLabelsEvaluateapplicationEvaluateApplicationPostResponse: *NewPolicyLabelsEvaluateapplicationEvaluateApplicationPostResponse(),
    }
    return m
}
// CreatePolicyLabelsEvaluateapplicationEvaluateApplicationResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePolicyLabelsEvaluateapplicationEvaluateApplicationResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPolicyLabelsEvaluateapplicationEvaluateApplicationResponse(), nil
}
// Deprecated: This class is obsolete. Use PolicyLabelsEvaluateapplicationEvaluateApplicationPostResponseable instead.
type PolicyLabelsEvaluateapplicationEvaluateApplicationResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PolicyLabelsEvaluateapplicationEvaluateApplicationPostResponseable
}
