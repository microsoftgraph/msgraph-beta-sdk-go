package privilegedapproval

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use MyrequestsMyRequestsGetResponseable instead.
type MyrequestsMyRequestsResponse struct {
    MyrequestsMyRequestsGetResponse
}
// NewMyrequestsMyRequestsResponse instantiates a new MyrequestsMyRequestsResponse and sets the default values.
func NewMyrequestsMyRequestsResponse()(*MyrequestsMyRequestsResponse) {
    m := &MyrequestsMyRequestsResponse{
        MyrequestsMyRequestsGetResponse: *NewMyrequestsMyRequestsGetResponse(),
    }
    return m
}
// CreateMyrequestsMyRequestsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMyrequestsMyRequestsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMyrequestsMyRequestsResponse(), nil
}
// Deprecated: This class is obsolete. Use MyrequestsMyRequestsGetResponseable instead.
type MyrequestsMyRequestsResponseable interface {
    MyrequestsMyRequestsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
