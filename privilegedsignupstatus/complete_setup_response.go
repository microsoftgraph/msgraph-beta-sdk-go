package privilegedsignupstatus

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CompleteSetupResponse 
// Deprecated: This class is obsolete. Use completeSetupPostResponse instead.
type CompleteSetupResponse struct {
    CompleteSetupPostResponse
}
// NewCompleteSetupResponse instantiates a new CompleteSetupResponse and sets the default values.
func NewCompleteSetupResponse()(*CompleteSetupResponse) {
    m := &CompleteSetupResponse{
        CompleteSetupPostResponse: *NewCompleteSetupPostResponse(),
    }
    return m
}
// CreateCompleteSetupResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCompleteSetupResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCompleteSetupResponse(), nil
}
// CompleteSetupResponseable 
// Deprecated: This class is obsolete. Use completeSetupPostResponse instead.
type CompleteSetupResponseable interface {
    CompleteSetupPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
