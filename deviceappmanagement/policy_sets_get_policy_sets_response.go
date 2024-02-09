package deviceappmanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type PolicySetsGetPolicySetsResponse struct {
    PolicySetsGetPolicySetsPostResponse
}
// NewPolicySetsGetPolicySetsResponse instantiates a new PolicySetsGetPolicySetsResponse and sets the default values.
func NewPolicySetsGetPolicySetsResponse()(*PolicySetsGetPolicySetsResponse) {
    m := &PolicySetsGetPolicySetsResponse{
        PolicySetsGetPolicySetsPostResponse: *NewPolicySetsGetPolicySetsPostResponse(),
    }
    return m
}
// CreatePolicySetsGetPolicySetsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePolicySetsGetPolicySetsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPolicySetsGetPolicySetsResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type PolicySetsGetPolicySetsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PolicySetsGetPolicySetsPostResponseable
}
