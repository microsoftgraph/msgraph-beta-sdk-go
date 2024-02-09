package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type GroupPolicyConfigurationsItemAssignResponse struct {
    GroupPolicyConfigurationsItemAssignPostResponse
}
// NewGroupPolicyConfigurationsItemAssignResponse instantiates a new GroupPolicyConfigurationsItemAssignResponse and sets the default values.
func NewGroupPolicyConfigurationsItemAssignResponse()(*GroupPolicyConfigurationsItemAssignResponse) {
    m := &GroupPolicyConfigurationsItemAssignResponse{
        GroupPolicyConfigurationsItemAssignPostResponse: *NewGroupPolicyConfigurationsItemAssignPostResponse(),
    }
    return m
}
// CreateGroupPolicyConfigurationsItemAssignResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGroupPolicyConfigurationsItemAssignResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyConfigurationsItemAssignResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type GroupPolicyConfigurationsItemAssignResponseable interface {
    GroupPolicyConfigurationsItemAssignPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
