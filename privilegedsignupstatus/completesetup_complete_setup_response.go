package privilegedsignupstatus

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use CompletesetupCompleteSetupPostResponseable instead.
type CompletesetupCompleteSetupResponse struct {
    CompletesetupCompleteSetupPostResponse
}
// NewCompletesetupCompleteSetupResponse instantiates a new CompletesetupCompleteSetupResponse and sets the default values.
func NewCompletesetupCompleteSetupResponse()(*CompletesetupCompleteSetupResponse) {
    m := &CompletesetupCompleteSetupResponse{
        CompletesetupCompleteSetupPostResponse: *NewCompletesetupCompleteSetupPostResponse(),
    }
    return m
}
// CreateCompletesetupCompleteSetupResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCompletesetupCompleteSetupResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCompletesetupCompleteSetupResponse(), nil
}
// Deprecated: This class is obsolete. Use CompletesetupCompleteSetupPostResponseable instead.
type CompletesetupCompleteSetupResponseable interface {
    CompletesetupCompleteSetupPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
