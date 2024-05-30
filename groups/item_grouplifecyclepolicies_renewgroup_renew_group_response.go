package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemGrouplifecyclepoliciesRenewgroupRenewGroupPostResponseable instead.
type ItemGrouplifecyclepoliciesRenewgroupRenewGroupResponse struct {
    ItemGrouplifecyclepoliciesRenewgroupRenewGroupPostResponse
}
// NewItemGrouplifecyclepoliciesRenewgroupRenewGroupResponse instantiates a new ItemGrouplifecyclepoliciesRenewgroupRenewGroupResponse and sets the default values.
func NewItemGrouplifecyclepoliciesRenewgroupRenewGroupResponse()(*ItemGrouplifecyclepoliciesRenewgroupRenewGroupResponse) {
    m := &ItemGrouplifecyclepoliciesRenewgroupRenewGroupResponse{
        ItemGrouplifecyclepoliciesRenewgroupRenewGroupPostResponse: *NewItemGrouplifecyclepoliciesRenewgroupRenewGroupPostResponse(),
    }
    return m
}
// CreateItemGrouplifecyclepoliciesRenewgroupRenewGroupResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemGrouplifecyclepoliciesRenewgroupRenewGroupResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemGrouplifecyclepoliciesRenewgroupRenewGroupResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemGrouplifecyclepoliciesRenewgroupRenewGroupPostResponseable instead.
type ItemGrouplifecyclepoliciesRenewgroupRenewGroupResponseable interface {
    ItemGrouplifecyclepoliciesRenewgroupRenewGroupPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
