package informationprotection

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PolicyLabelsEvaluateRemovalResponse 
// Deprecated: This class is obsolete. Use evaluateRemovalPostResponse instead.
type PolicyLabelsEvaluateRemovalResponse struct {
    PolicyLabelsEvaluateRemovalPostResponse
}
// NewPolicyLabelsEvaluateRemovalResponse instantiates a new PolicyLabelsEvaluateRemovalResponse and sets the default values.
func NewPolicyLabelsEvaluateRemovalResponse()(*PolicyLabelsEvaluateRemovalResponse) {
    m := &PolicyLabelsEvaluateRemovalResponse{
        PolicyLabelsEvaluateRemovalPostResponse: *NewPolicyLabelsEvaluateRemovalPostResponse(),
    }
    return m
}
// CreatePolicyLabelsEvaluateRemovalResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePolicyLabelsEvaluateRemovalResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPolicyLabelsEvaluateRemovalResponse(), nil
}
// PolicyLabelsEvaluateRemovalResponseable 
// Deprecated: This class is obsolete. Use evaluateRemovalPostResponse instead.
type PolicyLabelsEvaluateRemovalResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PolicyLabelsEvaluateRemovalPostResponseable
}
