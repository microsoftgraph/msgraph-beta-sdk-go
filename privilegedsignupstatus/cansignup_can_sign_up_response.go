package privilegedsignupstatus

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use CansignupCanSignUpGetResponseable instead.
type CansignupCanSignUpResponse struct {
    CansignupCanSignUpGetResponse
}
// NewCansignupCanSignUpResponse instantiates a new CansignupCanSignUpResponse and sets the default values.
func NewCansignupCanSignUpResponse()(*CansignupCanSignUpResponse) {
    m := &CansignupCanSignUpResponse{
        CansignupCanSignUpGetResponse: *NewCansignupCanSignUpGetResponse(),
    }
    return m
}
// CreateCansignupCanSignUpResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCansignupCanSignUpResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCansignupCanSignUpResponse(), nil
}
// Deprecated: This class is obsolete. Use CansignupCanSignUpGetResponseable instead.
type CansignupCanSignUpResponseable interface {
    CansignupCanSignUpGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
