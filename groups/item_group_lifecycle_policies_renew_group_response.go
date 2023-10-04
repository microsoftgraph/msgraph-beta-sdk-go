package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemGroupLifecyclePoliciesRenewGroupResponse 
// Deprecated: This class is obsolete. Use renewGroupPostResponse instead.
type ItemGroupLifecyclePoliciesRenewGroupResponse struct {
    ItemGroupLifecyclePoliciesRenewGroupPostResponse
}
// NewItemGroupLifecyclePoliciesRenewGroupResponse instantiates a new ItemGroupLifecyclePoliciesRenewGroupResponse and sets the default values.
func NewItemGroupLifecyclePoliciesRenewGroupResponse()(*ItemGroupLifecyclePoliciesRenewGroupResponse) {
    m := &ItemGroupLifecyclePoliciesRenewGroupResponse{
        ItemGroupLifecyclePoliciesRenewGroupPostResponse: *NewItemGroupLifecyclePoliciesRenewGroupPostResponse(),
    }
    return m
}
// CreateItemGroupLifecyclePoliciesRenewGroupResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemGroupLifecyclePoliciesRenewGroupResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemGroupLifecyclePoliciesRenewGroupResponse(), nil
}
// ItemGroupLifecyclePoliciesRenewGroupResponseable 
// Deprecated: This class is obsolete. Use renewGroupPostResponse instead.
type ItemGroupLifecyclePoliciesRenewGroupResponseable interface {
    ItemGroupLifecyclePoliciesRenewGroupPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
