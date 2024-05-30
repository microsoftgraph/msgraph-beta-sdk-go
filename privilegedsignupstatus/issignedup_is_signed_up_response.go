package privilegedsignupstatus

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use IssignedupIsSignedUpGetResponseable instead.
type IssignedupIsSignedUpResponse struct {
    IssignedupIsSignedUpGetResponse
}
// NewIssignedupIsSignedUpResponse instantiates a new IssignedupIsSignedUpResponse and sets the default values.
func NewIssignedupIsSignedUpResponse()(*IssignedupIsSignedUpResponse) {
    m := &IssignedupIsSignedUpResponse{
        IssignedupIsSignedUpGetResponse: *NewIssignedupIsSignedUpGetResponse(),
    }
    return m
}
// CreateIssignedupIsSignedUpResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIssignedupIsSignedUpResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIssignedupIsSignedUpResponse(), nil
}
// Deprecated: This class is obsolete. Use IssignedupIsSignedUpGetResponseable instead.
type IssignedupIsSignedUpResponseable interface {
    IssignedupIsSignedUpGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
