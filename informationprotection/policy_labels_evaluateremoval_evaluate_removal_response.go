package informationprotection

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use PolicyLabelsEvaluateremovalEvaluateRemovalPostResponseable instead.
type PolicyLabelsEvaluateremovalEvaluateRemovalResponse struct {
    PolicyLabelsEvaluateremovalEvaluateRemovalPostResponse
}
// NewPolicyLabelsEvaluateremovalEvaluateRemovalResponse instantiates a new PolicyLabelsEvaluateremovalEvaluateRemovalResponse and sets the default values.
func NewPolicyLabelsEvaluateremovalEvaluateRemovalResponse()(*PolicyLabelsEvaluateremovalEvaluateRemovalResponse) {
    m := &PolicyLabelsEvaluateremovalEvaluateRemovalResponse{
        PolicyLabelsEvaluateremovalEvaluateRemovalPostResponse: *NewPolicyLabelsEvaluateremovalEvaluateRemovalPostResponse(),
    }
    return m
}
// CreatePolicyLabelsEvaluateremovalEvaluateRemovalResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePolicyLabelsEvaluateremovalEvaluateRemovalResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPolicyLabelsEvaluateremovalEvaluateRemovalResponse(), nil
}
// Deprecated: This class is obsolete. Use PolicyLabelsEvaluateremovalEvaluateRemovalPostResponseable instead.
type PolicyLabelsEvaluateremovalEvaluateRemovalResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PolicyLabelsEvaluateremovalEvaluateRemovalPostResponseable
}
