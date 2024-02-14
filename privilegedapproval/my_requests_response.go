package privilegedapproval

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type MyRequestsResponse struct {
    MyRequestsGetResponse
}
// NewMyRequestsResponse instantiates a new MyRequestsResponse and sets the default values.
func NewMyRequestsResponse()(*MyRequestsResponse) {
    m := &MyRequestsResponse{
        MyRequestsGetResponse: *NewMyRequestsGetResponse(),
    }
    return m
}
// CreateMyRequestsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMyRequestsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMyRequestsResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type MyRequestsResponseable interface {
    MyRequestsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
