package informationprotection

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PolicyLabelsEvaluateApplicationResponse 
// Deprecated: This class is obsolete. Use evaluateApplicationPostResponse instead.
type PolicyLabelsEvaluateApplicationResponse struct {
    PolicyLabelsEvaluateApplicationPostResponse
}
// NewPolicyLabelsEvaluateApplicationResponse instantiates a new PolicyLabelsEvaluateApplicationResponse and sets the default values.
func NewPolicyLabelsEvaluateApplicationResponse()(*PolicyLabelsEvaluateApplicationResponse) {
    m := &PolicyLabelsEvaluateApplicationResponse{
        PolicyLabelsEvaluateApplicationPostResponse: *NewPolicyLabelsEvaluateApplicationPostResponse(),
    }
    return m
}
// CreatePolicyLabelsEvaluateApplicationResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePolicyLabelsEvaluateApplicationResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPolicyLabelsEvaluateApplicationResponse(), nil
}
// PolicyLabelsEvaluateApplicationResponseable 
// Deprecated: This class is obsolete. Use evaluateApplicationPostResponse instead.
type PolicyLabelsEvaluateApplicationResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PolicyLabelsEvaluateApplicationPostResponseable
}
