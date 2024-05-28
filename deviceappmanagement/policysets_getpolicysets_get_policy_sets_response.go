package deviceappmanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use PolicysetsGetpolicysetsGetPolicySetsPostResponseable instead.
type PolicysetsGetpolicysetsGetPolicySetsResponse struct {
    PolicysetsGetpolicysetsGetPolicySetsPostResponse
}
// NewPolicysetsGetpolicysetsGetPolicySetsResponse instantiates a new PolicysetsGetpolicysetsGetPolicySetsResponse and sets the default values.
func NewPolicysetsGetpolicysetsGetPolicySetsResponse()(*PolicysetsGetpolicysetsGetPolicySetsResponse) {
    m := &PolicysetsGetpolicysetsGetPolicySetsResponse{
        PolicysetsGetpolicysetsGetPolicySetsPostResponse: *NewPolicysetsGetpolicysetsGetPolicySetsPostResponse(),
    }
    return m
}
// CreatePolicysetsGetpolicysetsGetPolicySetsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePolicysetsGetpolicysetsGetPolicySetsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPolicysetsGetpolicysetsGetPolicySetsResponse(), nil
}
// Deprecated: This class is obsolete. Use PolicysetsGetpolicysetsGetPolicySetsPostResponseable instead.
type PolicysetsGetpolicysetsGetPolicySetsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PolicysetsGetpolicysetsGetPolicySetsPostResponseable
}
