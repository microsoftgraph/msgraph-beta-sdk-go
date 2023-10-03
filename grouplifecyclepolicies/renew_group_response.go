package grouplifecyclepolicies

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RenewGroupResponse 
// Deprecated: This class is obsolete. Use renewGroupPostResponse instead.
type RenewGroupResponse struct {
    RenewGroupPostResponse
}
// NewRenewGroupResponse instantiates a new RenewGroupResponse and sets the default values.
func NewRenewGroupResponse()(*RenewGroupResponse) {
    m := &RenewGroupResponse{
        RenewGroupPostResponse: *NewRenewGroupPostResponse(),
    }
    return m
}
// CreateRenewGroupResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRenewGroupResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRenewGroupResponse(), nil
}
// RenewGroupResponseable 
// Deprecated: This class is obsolete. Use renewGroupPostResponse instead.
type RenewGroupResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RenewGroupPostResponseable
}
