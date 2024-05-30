package grouplifecyclepolicies

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use RenewgroupRenewGroupPostResponseable instead.
type RenewgroupRenewGroupResponse struct {
    RenewgroupRenewGroupPostResponse
}
// NewRenewgroupRenewGroupResponse instantiates a new RenewgroupRenewGroupResponse and sets the default values.
func NewRenewgroupRenewGroupResponse()(*RenewgroupRenewGroupResponse) {
    m := &RenewgroupRenewGroupResponse{
        RenewgroupRenewGroupPostResponse: *NewRenewgroupRenewGroupPostResponse(),
    }
    return m
}
// CreateRenewgroupRenewGroupResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRenewgroupRenewGroupResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRenewgroupRenewGroupResponse(), nil
}
// Deprecated: This class is obsolete. Use RenewgroupRenewGroupPostResponseable instead.
type RenewgroupRenewGroupResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RenewgroupRenewGroupPostResponseable
}
